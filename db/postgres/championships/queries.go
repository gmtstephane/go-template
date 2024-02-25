package championships

import "github.com/Masterminds/squirrel"

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var (
	base = sq.
		Select(
			"championship.id",
			"championship.name",
			"championship.sport_id",
			"championship.icon_svg",
			"championship.created_at",
			"championship.updated_at",
		).
		From("championship")
)
