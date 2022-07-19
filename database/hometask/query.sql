CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(100),
    password TEXT
);