package vote

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func WriterHandler(vw Writer) echo.HandlerFunc {
    return func(c echo.Context) error {
        var v Vote
        if err := c.Bind(&v); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid vote data")
        }

        // Get email from context (assuming middleware sets it)
        email, ok := c.Get("email").(string)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
        }
        v.Email = email

        var result struct {
            ID uuid.UUID `json:"id"`
        }
        var err error

        result.ID, err = vw.Write(c.Request().Context(), &v)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to write vote")
        }

        return c.JSON(http.StatusCreated, result)
    }
}
