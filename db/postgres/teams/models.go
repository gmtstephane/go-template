package teams

import (
	"time"

	"github.com/gmtstephane/go-template/models"
)

type team struct {
	Id           int     `db:"id"`
	Name         string  `db:"name"`
	IconSvg      string  `db:"icon_svg"`
	SportID      int     `db:"sport_id"`
	SportName    string  `db:"sport_name"`
	SportType    string  `db:"sport_type"`
	LocationID   int     `db:"location_id"`
	LocationName string  `db:"location_name"`
	Address      string  `db:"address"`
	City         string  `db:"city"`
	State        string  `db:"state"`
	Country      string  `db:"country"`
	Zipcode      int     `db:"zipcode"`
	Latitude     float64 `db:"latitude"`
	Longitude    float64 `db:"longitude"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type teams []team

func (l teams) toModel() []models.Team {
	teams := make([]models.Team, len(l))
	for i, loc := range l {
		teams[i] = loc.toModel()
	}
	return teams
}

func (l team) toModel() models.Team {
	return models.Team{
		ID:   l.Id,
		Name: l.Name,
		Sport: models.Sport{
			ID:   l.SportID,
			Name: l.SportName,
			Type: models.SportType(l.SportType),
		},
		Icon: l.IconSvg,
		Location: models.Location{
			ID:      l.LocationID,
			Name:    l.LocationName,
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
		},
		UpdatedAt: l.UpdatedAt,
		CreatedAt: l.CreatedAt,
	}
}
