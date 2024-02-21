// package app handles http routes for html interaction
// either GET to render HTML pages or POST/PUT/DELETE with form elements
package app

import (
	"embed"
	"net/http"

	"github.com/gmtstephane/go-template/routes/app/middlewares"
	"github.com/labstack/echo/v4"
)

//go:embed public
var public embed.FS

type AppStorer interface {
	LocationStorer
}

func App(e *echo.Echo, s AppStorer) {
	e.GET("/public/*", echo.WrapHandler(http.FileServer(http.FS(public))))

	admingroup := e.Group("/admin", middlewares.WithUser)

	locations(admingroup, s)

}
