package middlewares

import (
	"context"
	"log/slog"
	"strings"

	"github.com/gmtstephane/go-template/internal/auth"
	"github.com/gmtstephane/go-template/models"
	"github.com/gmtstephane/go-template/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func WithUser(next echo.HandlerFunc) echo.HandlerFunc {

	//TODO : implement user authentication
	return func(c echo.Context) error {
		u := models.User{
			Picture: "https://lh3.googleusercontent.com/ogw/ANLem4YPUfUlP6R1Vv1D9Bx7Z_E-jnyP7lwbbMV_xOCecy0=s64-c-mo",
		}
		c.SetRequest(c.Request().WithContext(auth.Context(c.Request().Context(), u)))
		return next(c)
	}

}

func RequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogMethod:   true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
				)
			} else {
				slog.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	})
}

// func Path() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			ctx := context.WithValue(c.Request().Context(), views.PathCtx{}, c.Request().URL.Path)
// 			c.SetRequest(c.Request().WithContext(ctx))
// 			return next(c)
// 		}
// 	}
// }

func HTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if hx := c.Request().Header.Get("Hx-Request"); hx != "" {
			ctx := context.WithValue(c.Request().Context(), views.HtmxRequestCtx{}, true)
			c.SetRequest(c.Request().WithContext(ctx))
		}

		if hx := c.Request().Header.Get("Hx-Trigger"); strings.Contains(hx, "filter") || strings.Contains(hx, "paginate") {
			ctx := context.WithValue(c.Request().Context(), views.HtmxFiltertCtx{}, true)
			c.SetRequest(c.Request().WithContext(ctx))
		}

		return next(c)
	}
}
