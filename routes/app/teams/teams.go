package teams

import (
	"context"
	"io"
	"path"
	"strconv"

	"github.com/gmtstephane/go-template/models"
	. "github.com/gmtstephane/go-template/routes/app/utils"
	"github.com/gmtstephane/go-template/views/components/form"
	view "github.com/gmtstephane/go-template/views/teams"
	"github.com/labstack/echo/v4"
)

type (
	Storer = GenericStorer[models.Team, models.NewTeam, models.SearchTeam]

	LocationSearcher interface {
		Search(ctx context.Context, search models.SearchLocation) ([]models.Location, error)
	}

	SportSearcher interface {
		Search(ctx context.Context, search models.SearchSport) ([]models.Sport, error)
	}

	handler struct {
		teamStorer  Storer
		sportStorer SportSearcher
		locStorer   LocationSearcher
	}
)

func New(ts Storer, ss SportSearcher, ls LocationSearcher) handler {
	return handler{
		sportStorer: ss,
		locStorer:   ls,
		teamStorer:  ts,
	}
}

type Search struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}

// List returns a list of items
func (l handler) List(c echo.Context) error {

	search, err := Bind[Search](c)
	if err != nil {
		return err
	}

	if len(search.Name) < 2 {
		search.Name = ""
	}

	data, err := l.teamStorer.Search(c.Request().Context(), models.SearchTeam{
		Name: search.Name,
		Pagination: models.Pagination{
			Size: 10,
			Page: search.Page,
		},
	})

	if err != nil {
		return err
	}
	return Render(c, view.List(data, search.Page))
}

// New returns a form to create a new item
func (f handler) New(c echo.Context) error {

	sports, err := f.sportStorer.Search(c.Request().Context(), models.SearchSport{})
	if err != nil {
		return err
	}
	sportsKV := []form.KV{}
	for _, s := range sports {
		sportsKV = append(sportsKV, form.KV{Key: s.Name, Value: strconv.Itoa(s.ID)})
	}

	locations, err := f.locStorer.Search(c.Request().Context(), models.SearchLocation{})
	if err != nil {
		return err
	}
	locationsKV := []form.KV{}
	for _, s := range locations {
		locationsKV = append(locationsKV, form.KV{Key: s.Name, Value: strconv.Itoa(s.ID)})
	}

	return Render(c, view.Form(view.TeamFormProps{
		Locations: locationsKV,
		Sports:    sportsKV,
	}))
}

// Edit returns a form to edit an item
func (f handler) Edit(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))

	// teams, err := f.teamStorer.Get(c.Request().Context(), id)
	// if err != nil {
	// 	return err
	// }

	sports, err := f.sportStorer.Search(c.Request().Context(), models.SearchSport{})
	if err != nil {
		return err
	}
	sportsKV := []form.KV{}
	for _, s := range sports {
		sportsKV = append(sportsKV, form.KV{Key: s.Name, Value: strconv.Itoa(s.ID)})
	}

	locations, err := f.locStorer.Search(c.Request().Context(), models.SearchLocation{})
	if err != nil {
		return err
	}
	locationsKV := []form.KV{}
	for _, s := range locations {
		locationsKV = append(locationsKV, form.KV{Key: s.Name, Value: strconv.Itoa(s.ID)})
	}

	return Render(c, view.Form(view.TeamFormProps{
		Locations: locationsKV,
		Sports:    sportsKV,
	}))

}

// POST creates a new item
func (l handler) POST(c echo.Context) error {

	req, err := Bind[view.New](c)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file-upload")
	if err != nil {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// // Read the file
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	req.Icon = string(data)

	new, err := l.teamStorer.Create(c.Request().Context(), req.ToModel())
	if err != nil {
		return err
	}

	return Render(c,
		form.Feedback{
			Json: new,
			Link: path.Join(path.Dir(c.Request().URL.Path), strconv.Itoa(new)),
		})
}

// PUT updates an item
func (f handler) PUT(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	req, err := Bind[view.New](c)
	if err != nil {
		return err
	}

	updated, err := f.teamStorer.Update(c.Request().Context(), id, req.ToModel())
	if err != nil {
		return err
	}

	return Render(c, form.Feedback{
		Json: updated,
		Link: c.Request().URL.Path,
	})
}

// DELETE removes an item
func (f handler) DELETE(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
