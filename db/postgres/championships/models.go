package championships

import (
	"time"

	"github.com/gmtstephane/go-template/models"
)

type championship struct {
	Id        int        `db:"id"`
	Name      string     `db:"name"`
	SportID   int        `db:"sport_id"`
	IconSvg   *string    `db:"icon_svg"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type championships []championship

func (l championships) toModel() []models.Championship {
	t := make([]models.Championship, len(l))
	for i, it := range l {
		t[i] = it.toModel()
	}
	return t
}

func (l championship) toModel() models.Championship {
	return models.Championship{
		ID:        l.Id,
		Name:      l.Name,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		Sport: models.Sport{
			ID: l.SportID,
		},
	}
}
