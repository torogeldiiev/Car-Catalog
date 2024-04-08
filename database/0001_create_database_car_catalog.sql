-- +goose Up
-- +goose StatementBegin
-- +goose Up
-- +goose StatementBegin
CREATE DATABASE car_catalog1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- As there's no safe way to drop a database in a cross-platform manner,
-- the down migration typically does not include dropping the database.
-- If needed, manual intervention or a separate script can be used to drop the database.
-- +goose StatementEnd

-- +goose StatementEnd
