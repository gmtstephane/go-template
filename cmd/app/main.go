package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gmtstephane/go-template/db/postgres/championships"
	"github.com/gmtstephane/go-template/db/postgres/locations"
	"github.com/gmtstephane/go-template/db/postgres/sports"
	"github.com/gmtstephane/go-template/db/postgres/teams"
	"github.com/gmtstephane/go-template/routes/app"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}
	var conn = sqlx.MustConnect("pgx", os.Getenv("DATABASE_URL_GO"))

	e := echo.New()
	app.App(e, app.AppStorer{
		Locations:     locations.NewRepository(conn),
		Teams:         teams.NewRepository(conn),
		Sports:        sports.NewRepository(conn),
		Championships: championships.NewRepository(conn),
	})

	slog.Info("Server started on port 8080")
	if err := http.ListenAndServe(":8080", e); err != nil {
		slog.Error(err.Error())
	}
}
