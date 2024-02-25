package models

import "time"

type Championship struct {
	ID        int
	Name      string
	Sport     Sport
	Teams     []Team
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type SearchChampionship struct {
	Name       string
	Pagination Pagination
	SportID    int
}
