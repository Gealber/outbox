# Outbox

Contains a simple silly example of the [Outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html) in action, using
Golang(not Java ;) ) as the language. This implementation is meant to be run the following these steps:

There are two versions for the implementation of the pattern, 

1. **v0.0.1** using [Polling Publisher](https://microservices.io/patterns/data/polling-publisher.html) pattern for Message Relay.
2. **v0.0.2** using [Transaction log tailing](https://microservices.io/patterns/data/transaction-log-tailing.html) pattern for Message Relay.

Useful links to give you more context about why this repo, take a look at these articles I wrote in my super blog.

1. [Implementing Transactional outbox pattern in Go](https://www.gealber.com/tx-outbox-pattern).
2. [Transactional outbox. Part II, Message Relay](https://www.gealber.com/msg-realay-tx-outbox-part-ii).
3. [Transaction log tailing on CockroachDB. Part III.](https://www.gealber.com/msg-realay-cockroach-tx-outbox-part-iii).

## Set up for v0.0.1

### Prepare .env file

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

### Spin up docker compose

We are using Postgres docker image for our purpose, because I personally don't want you to install Postgres on your PC.

```bash

docker-compose --env-file .env up
```

### Run migrations

For this purpose you need to have [goose](https://github.com/pressly/goose), or run the SQL commands yourself. These commands are in **database/migrations**
folder.

Open another terminal and run...

```bash

make migration-up

```


## Setup for v0.0.2

### Prepare .env file

```bash

# App
APP_NAME=outbox
APP_ENV=development
APP_DEBUG=true
APP_PORT=8081

# Gin
# Gin modes: debug release test
GIN_MODE=debug

# Cockroach
COCKROACH_DSN="<COCKROACH_DSN>"

# Google cloud platform project id where the pub/sub is configured.
GCP_PROJECT_ID="<PROJECT_ID>"

```

### Spinning up cockroach.


```bash
cockroach start-single-node --insecure

```

After this command in the logs you will see the cockroach DSN, under de name `sql`. Keep that in mind.

You will need to open another terminal and run

```bash

cockroach sql --url="<DSN>" --format=csv

```


### Run migrations in Cockroach


```bash
make migration-cockroach-up
```

Always double check inside the Makefile, if the variable DB_CS is well formed, specially the port. This should be the value cockroach DSN.

##  Run the application

Just run `make run`


Good luck with this example :)...no one will see this Gealber.


