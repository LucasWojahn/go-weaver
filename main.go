package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/LucasWojahn/go-weaver/auth"
	"github.com/LucasWojahn/go-weaver/feedback"
	"github.com/LucasWojahn/go-weaver/vote"
	"github.com/ServiceWeaver/weaver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	feedback weaver.Ref[feedback.Writer]
	vote     weaver.Ref[vote.Writer]
	auth     weaver.Ref[auth.Auth]
	api      weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	fdb := app.feedback.Get()
	vt := app.vote.Get()
	us := app.auth.Get()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/health", auth.HealthHandler(us))
	e.POST("/auth", auth.Handler(us))

	// Auth Middleware for Echo
	authMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			email, err := us.ValidateToken(c.Request().Context(), c.Request().Header.Get("Authorization"))
			if err != nil {
				return c.NoContent(http.StatusForbidden) // Use echo's NoContent for error handling
			}
			c.Set("email", email) // Set the email in the Echo context
			return next(c)
		}
	}

	feedbackHandler := feedback.WriteHandler(fdb)
	voteHandler := vote.WriterHandler(vt)

	// Feedback Group
	// feedbackGroup := e.Group("/feedback")
	// feedbackGroup.Use(authMiddleware)
	e.POST("/feedback", feedbackHandler, authMiddleware)

	// Vote Group
	// voteGroup := e.Group("/vote")
	// voteGroup.Use(authMiddleware)
	e.POST("/vote", voteHandler, authMiddleware)

	fmt.Printf("Listener available on %v\n", app.api)
	e.Listener = app.api

	return e.Start("")
}
