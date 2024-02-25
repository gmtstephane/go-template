package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/gmtstephane/go-template/db/seed"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

// fefine flags for drop or up
var dropDB = flag.Bool("drop", false, "drop the database")
var seedDB = flag.Bool("seed", false, "seed the database")

func main() {

	flag.Parse()

	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}

	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_URL_GO"))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres", driver)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	if *seedDB {
		slog.Info("Seeding the database")
		seed.Ligue1(db)
		return
	}

	if *dropDB {
		slog.Info("Dropping the database")
		err = m.Drop()
	} else {
		slog.Info("Migrating the database")
		err = m.Up()
	}

	if err != nil {
		slog.Error(err.Error())
		return
	}

}
