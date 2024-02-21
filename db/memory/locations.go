package memory

import (
	"context"
	"time"

	"github.com/gmtstephane/go-template/db"
	"github.com/gmtstephane/go-template/models"
)

// locationDao is the data access object for the location model
type locationDao struct {
	ID        int
	Name      string
	Address   string
	City      string
	State     string
	Country   string
	Zipcode   int
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l locationDao) toModel() models.Location {
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
		CreatedAt:        l.CreatedAt,
		UpdatedAt:        l.UpdatedAt,
		LocationUserInfo: nil,
	}
}

func (r *Repository) SearchLocations(ctx context.Context) ([]models.Location, error) {
	locations := make([]models.Location, 0, len(r.store))
	for _, l := range r.store {
		locations = append(locations, l.toModel())
	}
	return locations, nil
}

func (r *Repository) CreateLocation(ctx context.Context, l models.NewLocation) {
	id := r.NewID()
	r.store[id] = locationDao{
		ID:        id,
		Name:      l.Name,
		Address:   l.Address,
		City:      l.City,
		State:     l.State,
		Country:   l.Country,
		Zipcode:   l.Zipcode,
		Latitude:  l.Geog.Latitude,
		Longitude: l.Geog.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *Repository) GetLocationByID(ctx context.Context, id int) (models.Location, error) {
	l, ok := r.store[id]
	if !ok {
		return models.Location{}, db.ErrNotFound
	}
	return l.toModel(), nil

}

func (r *Repository) UpdateLocation(ctx context.Context, id int, l models.NewLocation) {
	r.store[id] = locationDao{
		ID:        id,
		Name:      l.Name,
		Address:   l.Address,
		City:      l.City,
		State:     l.State,
		Country:   l.Country,
		Zipcode:   l.Zipcode,
		Latitude:  l.Geog.Latitude,
		Longitude: l.Geog.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *Repository) DeleteLocation(ctx context.Context, id int) {
	delete(r.store, id)
}
