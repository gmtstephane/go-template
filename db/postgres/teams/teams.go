package teams

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

func (p Repository) Create(ctx context.Context, newLoc models.NewTeam) (id int, err error) {
	return id, sq.
		Insert("team").
		Columns("name", "location_id", "sport_id", "icon_svg").
		Values(newLoc.Name, newLoc.LocationID, newLoc.SportID, newLoc.Icon).
		Suffix("RETURNING id").
		RunWith(p.db).QueryRow().Scan(&id)
}

func (p Repository) Delete(id int) error {
	// Define the SQL DELETE statement
	query := `
			DELETE FROM team WHERE id = $1
	`

	// Execute the SQL DELETE statement
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Return nil error
	return nil
}

func (p Repository) Get(ctx context.Context, id int) (models.Team, error) {
	t := team{}
	if err := p.db.QueryRowxContext(ctx, getByID, id).StructScan(&t); err != nil {
		return models.Team{}, err
	}
	return t.toModel(), nil
}

// SearchLocations searches for locations based on the provided search criteria.
// It returns a list of ticketerie.Location objects and an error, if any.
func (p Repository) Search(ctx context.Context, search models.SearchTeam) ([]models.Team, error) {
	var query = base

	if search.Pagination.Size > 0 {
		query = query.Limit(uint64(search.Pagination.Size)).Offset(uint64((search.Pagination.Page) * search.Pagination.Size))
	}

	if search.Name != "" {
		query = query.Where("unaccent(team.name) ILIKE unaccent(?)", "%"+search.Name+"%")
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var t teams
	err = p.db.Select(&t, sql, args...)
	if err != nil {
		return nil, err
	}

	return t.toModel(), nil
}

func (p Repository) Update(ctx context.Context, id int, update models.NewTeam) (models.Team, error) {
	panic("not implented")
}
