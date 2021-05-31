CREATE TABLE IF NOT EXISTS books
(
    id          serial PRIMARY KEY,
    isbn        varchar(255) UNIQUE,
    title       varchar(255),
    author      varchar(70),
    description text
)