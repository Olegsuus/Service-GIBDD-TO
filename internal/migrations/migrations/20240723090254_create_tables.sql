-- +goose Up
CREATE TABLE automobiles (
     id SERIAL PRIMARY KEY,
     release_date DATE NOT NULL,
     model VARCHAR(100) NOT NULL,
     license_plate VARCHAR(10) UNIQUE NOT NULL,
     registration_date DATE NOT NULL
);

CREATE TABLE inspections (
    id SERIAL PRIMARY KEY,
    automobile_id INTEGER NOT NULL,
    card_number VARCHAR(20) UNIQUE NOT NULL,
    inspection_date DATE NOT NULL,
    notes TEXT,
    FOREIGN KEY (automobile_id) REFERENCES automobiles(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS inspections;
DROP TABLE IF EXISTS automobiles;