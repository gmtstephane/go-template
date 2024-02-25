package sports

import (
	"time"

	"github.com/gmtstephane/go-template/models"
)

type sport struct {
	Id        int        `db:"id"`
	Name      string     `db:"name"`
	Type      string     `db:"type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type sports []sport

func (l sports) toModel() []models.Sport {
	t := make([]models.Sport, len(l))
	for i, loc := range l {
		t[i] = loc.toModel()
	}
	return t
}

func (l sport) toModel() models.Sport {
	return models.Sport{
		ID:        l.Id,
		Name:      l.Name,
		Type:      models.SportType(l.Type),
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}
