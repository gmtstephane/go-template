// package app handles http routes for html interaction
// either GET to render HTML pages or POST/PUT/DELETE with form elements
package app

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gmtstephane/go-template/routes/app/championships"
	"github.com/gmtstephane/go-template/routes/app/locations"
	"github.com/gmtstephane/go-template/routes/app/middlewares"
	"github.com/gmtstephane/go-template/routes/app/sports"
	"github.com/gmtstephane/go-template/routes/app/teams"
	"github.com/gmtstephane/go-template/views"
	"github.com/labstack/echo/v4"
)

//go:embed public
var public embed.FS

type AppStorer struct {
	Locations     locations.Storer
	Teams         teams.Storer
	Sports        sports.Storer
	Championships championships.Storer
}

func App(e *echo.Echo, s AppStorer) {
	e.Use(middlewares.HTMX)
	Public(e)
	Database(e.Group(views.DatabaseRoute, middlewares.WithUser, middlewares.RequestLogger()), s)

	for _, r := range e.Routes() {
		fmt.Println(r.Method, r.Path, r.Name)
	}
}

func Public(e *echo.Echo) {
	e.GET("/public/*", echo.WrapHandler(http.FileServer(http.FS(public))))
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/database/locations")
	})
}

func Database(e *echo.Group, s AppStorer) {
	locGroup := e.Group(views.LocationsRoute)
	locHandler := locations.New(s.Locations)
	locGroup.GET("/components/select", locHandler.Select)
	Register(locGroup, locHandler)

	teamGroup := e.Group(views.TeamsRoute)
	teamHandler := teams.New(s.Teams, s.Sports, s.Locations)
	Register(teamGroup, teamHandler)

	sportGroup := e.Group(views.SportsRoute)
	sprtHandler := sports.New(s.Sports)
	sportGroup.GET("/components/select", sprtHandler.Select)

	championshipGroup := e.Group(views.ChampionshipsRoute)
	championshipHandler := championships.New(s.Championships)
	championshipGroup.GET("/components/select", championshipHandler.Select)
}
