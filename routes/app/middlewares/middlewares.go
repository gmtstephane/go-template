package middlewares

import (
	"github.com/gmtstephane/go-template/internal/auth"
	"github.com/gmtstephane/go-template/models"
	"github.com/labstack/echo/v4"
)

func WithUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := models.User{
			Picture: "https://lh3.googleusercontent.com/ogw/ANLem4YPUfUlP6R1Vv1D9Bx7Z_E-jnyP7lwbbMV_xOCecy0=s64-c-mo",
		}
		c.SetRequest(c.Request().WithContext(auth.Context(c.Request().Context(), u)))
		return next(c)
	}
}
