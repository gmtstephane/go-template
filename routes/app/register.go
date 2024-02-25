package app

import "github.com/labstack/echo/v4"

type ServerRessource interface {
	//List returns a list of items
	List(c echo.Context) error
	//New returns a form to create a new item
	New(c echo.Context) error
	//Edit returns a form to edit an item
	Edit(c echo.Context) error

	//POST creates a new item
	POST(c echo.Context) error
	//PUT updates an item
	PUT(c echo.Context) error
	//DELETE removes an item
	DELETE(c echo.Context) error
}

func Register(e *echo.Group, r ServerRessource, m ...echo.MiddlewareFunc) {

	e.GET("", r.List)

	e.GET("/new", r.New)
	e.POST("/new", r.POST)

	e.GET("/:id", r.Edit)
	e.PUT("/:id", r.PUT)
	e.DELETE("/:id", r.DELETE)
}
