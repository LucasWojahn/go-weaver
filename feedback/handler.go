package feedback

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// WriteHandler handle the write feedback request
func WriteHandler(fw Writer) echo.HandlerFunc {
	return func(c echo.Context) error {
		var f Feedback
		if err := c.Bind(&f); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid feedback data")
		}

		// Get email from context (assuming middleware sets it)
		email, ok := c.Get("email").(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		f.Email = email

		var result struct {
			ID uuid.UUID `json:"id"`
		}

		var err error

		result.ID, err = fw.Write(c.Request().Context(), &f)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to write feedback")
		}

		return c.JSON(http.StatusCreated, result)
	}
}
