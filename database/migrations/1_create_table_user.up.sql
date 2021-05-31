CREATE TABLE users
(
    id         serial PRIMARY KEY,
    name       VARCHAR(255),
    password   VARCHAR(255),
    email      VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)