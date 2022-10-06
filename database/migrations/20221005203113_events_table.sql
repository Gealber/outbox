-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS events (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (), 
    type varchar(100),
    resource_id uuid,
    resource_type varchar(100),
    published boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
