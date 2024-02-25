package models

import (
	"strconv"
	"time"

	"github.com/gmtstephane/go-template/views/components/form"
)

type (
	Point struct {
		Latitude  float64
		Longitude float64
	}

	LocationUserInfo struct {
		DistanceFromUser float64
	}

	Location struct {
		ID        int
		Name      string
		Address   string
		City      string
		State     string
		Country   string
		Zipcode   int
		Geog      Point
		CreatedAt time.Time
		UpdatedAt *time.Time
		// LocationUserInfo *LocationUserInfo
	}

	Locations []Location

	NewLocation struct {
		Name    string
		Address string
		City    string
		State   string
		Country string
		Zipcode int
		Geog    Point
	}

	SearchLocation struct {
		DistanceSearch DistanceSearch
		Name           string
		Pagination     Pagination
	}

	Pagination struct {
		Page int
		Size int
	}

	DistanceSearch struct {
		Origin   Point
		Distance float64
	}
)

func (d DistanceSearch) IsValid() bool {
	return d.Origin.Latitude != 0 && d.Origin.Longitude != 0 && d.Distance >= 0
}

func (lcs Locations) KV() []form.KV {
	kvs := make([]form.KV, len(lcs))
	for i, lc := range lcs {
		kvs[i] = form.KV{Key: lc.Name, Value: strconv.Itoa(lc.ID)}
	}
	return kvs
}
