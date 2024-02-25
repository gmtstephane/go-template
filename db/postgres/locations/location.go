package locations

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/gmtstephane/go-template/db"
	"github.com/gmtstephane/go-template/models"
	"github.com/jmoiron/sqlx"
)

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

const (
	queryWithLocation = `
WITH location_with_distance AS
	(
		SELECT 
			id,
			name,
			address,
			city,
			state,
			country,
			zipcode,
			latitude,
			longitude,
			created_at,
			(6371 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude)))) AS distance
		FROM	location
	)
`

	queryWithoutLocation = `
WITH location_with_distance AS 
	(
		SELECT id, name, address, city, state, country, zipcode, latitude, longitude, created_at, 0 AS distance
		FROM location
	)
`
)

type location struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Address   string     `db:"address"`
	City      string     `db:"city"`
	State     string     `db:"state"`
	Country   string     `db:"country"`
	Zipcode   int        `db:"zipcode"`
	Latitude  float64    `db:"latitude"`
	Longitude float64    `db:"longitude"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Distance  float64    `db:"distance"`
}

type locations []location

func (l locations) toModel() []models.Location {
	locations := make([]models.Location, len(l))
	for i, loc := range l {
		locations[i] = loc.toModel()
	}
	return locations
}

func (l location) toModel() models.Location {
	// var locuser *models.LocationUserInfo = nil
	// if l.Distance != 0 {
	// 	locuser = &models.LocationUserInfo{DistanceFromUser: l.Distance}
	// }
	return models.Location{
		ID:      l.ID,
		Name:    l.Name,
		Address: l.Address,
		City:    l.City,
		State:   l.State,
		Country: l.Country,
		Zipcode: l.Zipcode,
		Geog: models.Point{
			Latitude:  l.Latitude,
			Longitude: l.Longitude,
		},
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		// LocationUserInfo: locuser,
	}
}

func (p Repository) Create(ctx context.Context, newLoc models.NewLocation) (int, error) {
	query := `
			INSERT INTO location (
					name, address, city, state, country, zipcode,
					latitude, longitude, created_at
			) VALUES (
					$1, $2, $3, $4, $5, $6, $7, $8, $9
			) RETURNING id
	`

	// Execute the SQL INSERT statement
	var locationID int
	ts := time.Now()
	err := p.db.QueryRow(
		query,
		newLoc.Name,
		newLoc.Address,
		newLoc.City,
		newLoc.State,
		newLoc.Country,
		newLoc.Zipcode,
		newLoc.Geog.Latitude,
		newLoc.Geog.Longitude,
		ts,
	).Scan(&locationID)
	if err != nil {
		return 0, err
	}

	// Return the created ticketerie.Location and nil error
	return locationID, nil
}

func (p Repository) Delete(id int) error {
	// Define the SQL DELETE statement
	query := `
			DELETE FROM location WHERE id = $1
	`

	// Execute the SQL DELETE statement
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Return nil error
	return nil
}

func (p Repository) Get(ctx context.Context, id int) (models.Location, error) {
	// Define the SQL SELECT statement with explicit fields
	query := `
			SELECT id, name, address, city, state, country, zipcode, latitude, longitude, created_at, updated_at
			FROM location WHERE id = $1
	`

	// Execute the SQL SELECT statement
	var location location
	err := p.db.Get(&location, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Location{}, errors.Join(err, db.ErrNotFound)
		}
		return models.Location{}, err
	}

	// Return the retrieved location and nil error
	return location.toModel(), nil
}

// SearchLocations searches for locations based on the provided search criteria.
// It returns a list of ticketerie.Location objects and an error, if any.
func (p Repository) Search(ctx context.Context, search models.SearchLocation) ([]models.Location, error) {

	query := sq.Select(
		"id",
		"name",
		"address",
		"city",
		"state",
		"country",
		"zipcode",
		"latitude",
		"longitude",
		"created_at",
		"distance").
		From("location_with_distance")

	if search.DistanceSearch.IsValid() {
		query = query.PrefixExpr(
			squirrel.Expr(queryWithLocation, search.DistanceSearch.Origin.Latitude, search.DistanceSearch.Origin.Longitude, search.DistanceSearch.Origin.Latitude))
		if search.DistanceSearch.Distance > 0 {
			query = query.Where(squirrel.LtOrEq{"distance": search.DistanceSearch.Distance})
		}
	} else {
		query = query.Prefix(queryWithoutLocation)
	}

	if search.Name != "" {
		query = query.Where("unaccent(name) ILIKE unaccent(?)", "%"+search.Name+"%")
	}

	if search.Pagination.Size > 0 {
		query = query.Limit(uint64(search.Pagination.Size)).Offset(uint64((search.Pagination.Page) * search.Pagination.Size))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var locations locations
	err = p.db.Select(&locations, sql, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return locations.toModel(), nil
}

func (p Repository) Update(ctx context.Context, id int, update models.NewLocation) (models.Location, error) {
	// Define the SQL UPDATE statement
	query := `
			UPDATE location
			SET name = $1, address = $2, city = $3, state = $4, country = $5, zipcode = $6, latitude = $7, longitude = $8
			WHERE id = $9
			RETURNING *
	`
	loc := location{}
	// Execute the SQL UPDATE statement
	err := p.db.QueryRowxContext(ctx,
		query,
		update.Name,
		update.Address,
		update.City,
		update.State,
		update.Country,
		update.Zipcode,
		update.Geog.Latitude,
		update.Geog.Longitude,
		id,
	).StructScan(&loc)
	if err != nil {
		return models.Location{}, err
	}

	// Return the updated ticketerie.Location and nil error
	return loc.toModel(), nil
}
