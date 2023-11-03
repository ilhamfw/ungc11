CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    longitude NUMERIC,
    latitude NUMERIC,
    rating NUMERIC
);