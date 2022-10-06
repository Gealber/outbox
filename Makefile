.DEFAULT_GOAL := help
.PHONY : build

ENV?=development

APP := outbox

# Database connection string
DB_CS="user=postgres password=example dbname=$(APP) sslmode=disable"

run: ## Run code
	@go run main.go

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/${APP}

migration-status: ## Migartion status
	@goose -dir=./database/migrations postgres $(DB_CS) status

migration-up: ## Migartion up
	@goose -dir=./database/migrations postgres $(DB_CS) up

migration-down: ## Migartion down
	@goose -dir=./database/migrations postgres $(DB_CS) down

migration-create: ## Migartion create (migration-create init)
	@goose -dir=./database/migrations postgres $(DB_CS) $* sql

clean: ## Cleaning binary
	@rm -f bin/${APP}

help: ## Show commands availables
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

