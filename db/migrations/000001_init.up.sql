CREATE TABLE location (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (name <> ''),
    address VARCHAR(255) NOT NULL CHECK (address <> ''),
    city VARCHAR(100) NOT NULL CHECK (city <> ''),
    state VARCHAR(50) NOT NULL CHECK (state <> ''),
    country VARCHAR(100) NOT NULL CHECK (country <> ''),
    zipcode INTEGER NOT NULL CHECK (zipcode > 0),
    latitude NUMERIC(10, 6) NOT NULL CHECK (latitude >= -90 AND latitude <= 90),
    longitude NUMERIC(10, 6) NOT NULL CHECK (longitude >= -180 AND longitude <= 180),
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT null,
    CONSTRAINT unique_location UNIQUE (address, city, state, country, zipcode)
);

-- Migration down script seems to not delete this properly
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sport_type') THEN
        CREATE TYPE sport_type AS ENUM ('Team', 'Individual', 'Event');
    END IF;
END$$;


CREATE TABLE sport (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (name <> ''),
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT null,
    type sport_type NOT NULL,
    CONSTRAINT unique_sport UNIQUE (name)
);


-- Create Championship Table 
CREATE TABLE championship (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL CHECK (name <> ''),
    sport_id INTEGER NOT NULL,
    icon_svg TEXT, 
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT null,
    CONSTRAINT unique_championship UNIQUE (name, sport_id),
    FOREIGN KEY (sport_id) REFERENCES sport(id)
);

CREATE TABLE team (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sport_id INTEGER NOT NULL,
    location_id INTEGER NOT NULL,
    icon_svg TEXT, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT unique_team UNIQUE (name, sport_id),
    FOREIGN KEY (sport_id) REFERENCES sport(id),
    FOREIGN KEY (location_id) REFERENCES location(id)
);