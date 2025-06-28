CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS engine (
    id UUID PRIMARY KEY,
    displacement INT NOT NULL CHECK (displacement > 0 AND displacement <= 10000),
    no_of_cylinders INT NOT NULL CHECK (no_of_cylinders > 0 AND no_of_cylinders <= 16),
    car_range INT NOT NULL CHECK (car_range > 0 AND car_range <= 1000),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);