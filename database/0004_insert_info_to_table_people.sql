-- +goose Up
-- +goose StatementBegin
INSERT INTO people (name, surname, patronymic) VALUES
('Azamat', 'Torogeldiev', 'Duishenbekovich'),
('Tilek', 'Kasymaliev', 'Aidarov');
-- +goose StatementEnd
-- +goose Down