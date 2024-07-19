package auth

import (
	"net/http"

	"github.com/labstack/echo/v4" // For Echo framework
)

// HealthHandler handle the health request
func HealthHandler(us Auth) echo.HandlerFunc {
	return func(c echo.Context) error {
		msg, err := us.Health(c.Request().Context())
		if err != nil {
			// Handle the error, e.g.:
			return echo.NewHTTPError(http.StatusInternalServerError, "Health check failed")
		}
		return c.String(http.StatusOK, msg) // Send the health message as plain text
	}
}

// Handler handle the auth request
func Handler(us Auth) echo.HandlerFunc {
	return func(c echo.Context) error {
		var param struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Bind(&param); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := us.ValidateUser(c.Request().Context(), param.Email, param.Password); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}

		token, err := us.GenerateToken(c.Request().Context(), param.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}
