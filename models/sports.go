package models

import "time"

type SportType string

const (
	SportTypeTeam       SportType = "Team"
	SportTypeIndividual SportType = "Individual"
	SportTypeEvent      SportType = "Event"
)

type (
	Sport struct {
		ID        int
		Name      string
		Type      SportType
		CreatedAt time.Time
		UpdatedAt *time.Time
	}
)

type SearchSport struct {
	Name       string
	Pagination Pagination
}
