package sports

import (
	"context"

	"github.com/gmtstephane/go-template/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (p Repository) Search(ctx context.Context, search models.SearchSport) ([]models.Sport, error) {
	var query = base

	if search.Pagination.Size > 0 {
		query = query.Limit(uint64(search.Pagination.Size)).Offset(uint64((search.Pagination.Page) * search.Pagination.Size))
	}

	if search.Name != "" {
		query = query.Where("unaccent(name) ILIKE unaccent(?)", "%"+search.Name+"%")
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var t sports
	err = p.db.Select(&t, sql, args...)
	if err != nil {
		return nil, err
	}

	return t.toModel(), nil
}
