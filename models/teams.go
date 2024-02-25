package models

import "time"

type Team struct {
	ID            int
	Name          string
	Location      Location
	Sport         Sport
	Icon          string
	Championships []Championship
	UpdatedAt     *time.Time
	CreatedAt     time.Time
}

type NewTeam struct {
	Name          string
	LocationID    int
	SportID       int
	Icon          string
	Championships []int
}

type SearchTeam struct {
	Name       string
	Pagination Pagination
}
