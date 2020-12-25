CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    uuid       varchar     NOT NULL,
    name       varchar     NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NULL
);
