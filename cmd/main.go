// package cmd holds command lines utility
package main

import (
	"log/slog"
	"net/http"

	"github.com/gmtstephane/go-template/db/memory"
	"github.com/gmtstephane/go-template/routes/app"
	"github.com/labstack/echo/v4"
)

func main() {

	store := memory.NewRepository()

	e := echo.New()
	app.App(e, store)

	slog.Info("Server started on port 8080")
	if err := http.ListenAndServe(":8080", e); err != nil {
		slog.Error(err.Error())
	}
}
