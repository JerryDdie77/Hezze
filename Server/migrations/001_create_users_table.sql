-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(30) NOT NULL UNIQUE,
    first_name VARCHAR(30) NOT NULL,
    surname VARCHAR(30) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) UNIQUE,
    is_blocked BOOLEAN DEFAULT FALSE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;