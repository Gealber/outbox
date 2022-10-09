-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cats (
    id uuid UNIQUE DEFAULT gen_random_uuid (), 
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
