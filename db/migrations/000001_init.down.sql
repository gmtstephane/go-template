DROP TABLE IF EXISTS team;

-- Drop the championship table first due to its dependencies
DROP TABLE IF EXISTS championship;

-- Drop the sport table next
DROP TABLE IF EXISTS sport;

-- Finally, drop the sport_type enum
DROP TYPE sport_type;

-- Drop the location table
DROP TABLE IF EXISTS location;


