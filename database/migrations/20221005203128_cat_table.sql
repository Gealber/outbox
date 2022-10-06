-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS cats (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (), 
    name varchar(100),
    color varchar(100),
    weight numeric(5,2),
    intelligence smallint,
    laziness smallint,
    curiosity smallint,
    sociability smallint,
    egoism smallint,
    miau_power smallint,
    attack smallint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cats;
-- +goose StatementEnd
