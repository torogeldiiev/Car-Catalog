-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down