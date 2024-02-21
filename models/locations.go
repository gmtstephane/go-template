package models

import "time"

type (
	Point struct {
		Latitude  float64
		Longitude float64
	}
	LocationUserInfo struct {
		DistanceFromUser float64
	}

	Location struct {
		ID               int
		Name             string
		Address          string
		City             string
		State            string
		Country          string
		Zipcode          int
		Geog             Point
		CreatedAt        time.Time
		UpdatedAt        time.Time
		LocationUserInfo *LocationUserInfo
	}

	NewLocation struct {
		Name    string
		Address string
		City    string
		State   string
		Country string
		Zipcode int
		Geog    Point
	}
)
