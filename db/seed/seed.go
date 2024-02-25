package seed

import (
	_ "embed"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

//go:embed data/ligue1.sql
var ligue1 string

func Ligue1(db *sqlx.DB) {
	_, err := db.Exec(ligue1)
	if err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("Ligue 1 data loaded")
	}
}
