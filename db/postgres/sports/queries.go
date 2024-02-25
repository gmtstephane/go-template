package sports

import "github.com/Masterminds/squirrel"

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var (
	base = sq.
		Select(
			"sport.id",
			"sport.name",
			"sport.type",
			"created_at",
			"updated_at",
		).
		From("sport")
)
