// package views holds html templates
package views

import (
	"context"
	"path"
)

const (
	DatabaseRoute      string = "/database"
	LocationsRoute     string = "/locations"
	SportRoute         string = "/sports"
	TeamsRoute         string = "/teams"
	SportsRoute        string = "/sports"
	ChampionshipsRoute string = "/championships"
)

func Path(r ...string) string {
	return path.Join(r...)
}

func PathNew(r ...string) string {
	return path.Join(append(r, "new")...)
}

// TrailNew removes /new at the end of the path
func Dir(s string) string {
	return path.Dir(s)
}

type HtmxRequestCtx struct{}
type HtmxFiltertCtx struct{}

func IsHTMX(ctx context.Context) bool {
	is, ok := ctx.Value(HtmxRequestCtx{}).(bool)
	if !ok {
		return false
	}
	return is
}

func IsHTMXFilter(ctx context.Context) bool {
	is, ok := ctx.Value(HtmxFiltertCtx{}).(bool)
	if !ok {
		return false
	}
	return is
}
