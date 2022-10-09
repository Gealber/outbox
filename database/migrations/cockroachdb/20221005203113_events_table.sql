-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id uuid UNIQUE DEFAULT gen_random_uuid (), 
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
