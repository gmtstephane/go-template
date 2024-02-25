package sports

import (
	"context"
	"strconv"

	"github.com/gmtstephane/go-template/models"
	. "github.com/gmtstephane/go-template/routes/app/utils"
	"github.com/gmtstephane/go-template/views/components/form"
	"github.com/labstack/echo/v4"
)

type (
	Storer interface {
		Search(ctx context.Context, search models.SearchSport) ([]models.Sport, error)
	}

	handler struct {
		Storer
	}
)

func New(ts Storer) handler {
	return handler{
		Storer: ts,
	}
}

type Search struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}

// List returns a list of items
func (l handler) List(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// New returns a form to create a new item
func (f handler) New(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Edit returns a form to edit an item
func (f handler) Edit(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// POST creates a new item
func (l handler) POST(c echo.Context) error {
	panic("not implemented") // TODO: Implement

}

// PUT updates an item
func (f handler) PUT(c echo.Context) error {
	panic("not implemented") // TODO: Implement

}

// DELETE removes an item
func (f handler) DELETE(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

type SearhSport struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}

func (f handler) Select(c echo.Context) error {

	data, err := f.Search(c.Request().Context(), models.SearchSport{})
	if err != nil {
		return err
	}

	componentVlues := []form.KV{}
	for _, loc := range data {
		componentVlues = append(componentVlues, form.KV{
			Value: strconv.Itoa(loc.ID),
			Key:   loc.Name,
		})
	}

	return Render(c, form.InputSelect("Sport", "sport", componentVlues).Width(form.InputWidthHalf))
}
