package teams

import "github.com/Masterminds/squirrel"

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var (
	base = sq.
		Select(
			"team.id",
			"team.name",
			"team.icon_svg",
			"team.created_at",
			"team.updated_at",
			"sport.id AS sport_id",
			"sport.name AS sport_name",
			"sport.type AS sport_type",
			"location.id AS location_id",
			"location.name AS location_name",
			"location.address",
			"location.city",
			"location.state",
			"location.country",
			"location.zipcode",
			"location.latitude",
			"location.longitude",
		).
		From("team").
		Join("sport ON team.sport_id = sport.id").
		Join("location ON team.location_id = location.id")

	getByID, _ = base.Where(squirrel.Eq{"team.id": "?"}).MustSql()
)
