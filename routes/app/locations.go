package app

import (
	"context"

	"github.com/gmtstephane/go-template/models"
	view "github.com/gmtstephane/go-template/views/locations"
	"github.com/labstack/echo/v4"
)

type (
	LocationStorer interface {
		SearchLocations(ctx context.Context) ([]models.Location, error)
	}

	locationHandler struct {
		storer LocationStorer
	}
)

func locations(e *echo.Group, s LocationStorer) {
	l := locationHandler{storer: s}
	e.GET("/locations", l.List)
	e.GET("/locations/new", l.New)
}

func (l locationHandler) List(c echo.Context) error {
	data, err := l.storer.SearchLocations(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.List(data))
}

func (l locationHandler) New(c echo.Context) error {
	return render(c, view.New())
}
