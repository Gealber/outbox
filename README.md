# Outbox

Contains a simple silly example of the [Outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html) in action, using
Golang(not Java ;) ) as the language. This implementation is meant to be run the following these steps:

## Prepare .env file

Create a `.env` file with the following content

```bash
# App
APP_NAME=outbox
APP_ENV=development
APP_DEBUG=true
APP_PORT=8080

# Gin
# Gin modes: debug release test
GIN_MODE=debug

# Database
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=example
DATABASE_HOST=0.0.0.0
DATABASE_NAME=outbox
DATABASE_TIMEZONE=UTC
DATABASE_SSLMODE=disable
DATABASE_PORT=5432
DATABASE_ENGINE=Postgres

```

## Spin up docker compose

We are using Postgres docker image for our purpose, because I personally don't want you to install Postgres on your PC.

```bash

docker-compose --env-file .env up
```

## Run migrations

For this purpose you need to have [goose](https://github.com/pressly/goose), or run the SQL commands yourself. These commands are in **database/migrations**
folder.

Open another terminal and run...

```bash

make migration-up

```

##  Run the application

Just run `make run`


Good luck with this example :)...no one will see this Gealber.


