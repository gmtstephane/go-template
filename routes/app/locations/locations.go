package locations

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gmtstephane/go-template/models"
	. "github.com/gmtstephane/go-template/routes/app/utils"
	"github.com/gmtstephane/go-template/views"
	"github.com/gmtstephane/go-template/views/components/form"
	view "github.com/gmtstephane/go-template/views/locations"
	"github.com/labstack/echo/v4"
)

type (
	Storer          = GenericStorer[models.Location, models.NewLocation, models.SearchLocation]
	locationHandler struct {
		store Storer
	}
)

func New(s Storer) locationHandler {
	return locationHandler{store: s}
}

type SearchLocation struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}

// List returns a list of items
func (l locationHandler) List(c echo.Context) error {

	locationParams, err := Bind[SearchLocation](c)
	if err != nil {
		return err
	}

	data, err := l.store.Search(c.Request().Context(), models.SearchLocation{
		Name: locationParams.Name,
		Pagination: models.Pagination{
			Size: 10,
			Page: locationParams.Page,
		},
	})

	if err != nil {
		return err
	}
	return Render(c, view.List(data, locationParams.Page))
}

// New returns a form to create a new item
func (f locationHandler) New(c echo.Context) error {
	return Render(c, view.Form())
}

// Edit returns a form to edit an item
func (f locationHandler) Edit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	loc, err := f.store.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return Render(c,
		view.Form().WithValues(view.NewLocation{
			Name:      loc.Name,
			Address:   loc.Address,
			City:      loc.City,
			State:     loc.State,
			Country:   loc.Country,
			Zipcode:   loc.Zipcode,
			Latitude:  loc.Geog.Latitude,
			Longitude: loc.Geog.Longitude,
		}).WithMethod(http.MethodPut))
}

// POST creates a new item
func (l locationHandler) POST(c echo.Context) error {

	fmt.Println("isHTMX: ", views.IsHTMX(c.Request().Context()))
	locForm, err := Bind[view.NewLocation](c)
	if err != nil {
		return err
	}

	id, err := l.store.Create(c.Request().Context(), locForm.ToModel())
	if err != nil {
		return err
	}

	return Render(c,
		form.Feedback{
			Json: id,
			Link: path.Join(path.Dir(c.Request().URL.Path), strconv.Itoa(id)),
		})
}

// PUT updates an item
func (f locationHandler) PUT(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	locForm, err := Bind[view.NewLocation](c)
	if err != nil {
		return err
	}

	loc, err := f.store.Update(c.Request().Context(), id, locForm.ToModel())
	if err != nil {
		return err
	}

	return Render(c, form.Feedback{
		Json: loc,
		Link: c.Request().URL.Path,
	})
}

// DELETE removes an item
func (f locationHandler) DELETE(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (f locationHandler) Select(c echo.Context) error {

	data, err := f.store.Search(c.Request().Context(), models.SearchLocation{})
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

	return Render(c, form.InputSelect("Location", "location", componentVlues).Width(form.InputWidthHalf))
}
