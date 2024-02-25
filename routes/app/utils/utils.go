package utils

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func Bind[T any](c echo.Context) (T, error) {
	v := new(T)
	if err := c.Bind(v); err != nil {
		return *v, err
	}
	return *v, nil
}

type GenericStorer[Model any, NewModel any, SearchModel any] interface {
	Search(ctx context.Context, s SearchModel) ([]Model, error)
	Create(ctx context.Context, n NewModel) (int, error)
	Get(ctx context.Context, id int) (Model, error)
	Update(ctx context.Context, id int, l NewModel) (Model, error)
}
