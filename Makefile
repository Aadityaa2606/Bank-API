include .env

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Database Management

startdb: ## Start PostgreSQL container
	sudo docker run --name $(POSTGRES_CONTAINER_NAME) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-p $(DB_PORT):$(DB_PORT) \
		-d postgres:17-alpine || sudo docker start $(POSTGRES_CONTAINER_NAME)

stopdb: ## Stop PostgreSQL container
	sudo docker stop $(POSTGRES_CONTAINER_NAME) || true

createdb: ## Creates the database
	sudo docker start $(POSTGRES_CONTAINER_NAME) || true
	sudo docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME) || \
	echo "Database already exists, skipping creation."

dropdb: ## Deletes the database
	sudo docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)

resetdb: ## Resets the database
	make dropdb
	make createdb

psql: ## Connect to PostgreSQL with psql
	sudo docker exec -it $(POSTGRES_CONTAINER_NAME) psql --username=$(DB_USER) --dbname=$(DB_NAME)

##@ Migrations

migrateup: ## Run all pending migrations
	goose -dir $(GOOSE_MIGRATION_DIR) postgres "$(DB_SOURCE)" up

migrateup1: ## Run 1 pending migration
	goose -dir $(GOOSE_MIGRATION_DIR) postgres "$(DB_SOURCE)" up 1

migratedown: ## Rollback all migrations
	goose -dir $(GOOSE_MIGRATION_DIR) postgres "$(DB_SOURCE)" down

migratedown1: ## Rollback 1 migration
	goose -dir $(GOOSE_MIGRATION_DIR) postgres "$(DB_SOURCE)" down 1

##@ Code Generation

sqlc: ## Generate Go code from SQL
	sqlc generate

proto: ## Generate Go code from Protocol Buffers
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto

##@ Testing and Running

test: ## Run tests with coverage (no cache)
	go test -v -cover -count=1 ./...

server: ## Run the server
	go run main.go

build: ## Build the application
	go build -o bin/simple_bank main.go 

start: ## Starts every service
	sudo systemctl start docker
	sleep 2
	make startdb
	sleep 5
	make createdb
	make migrateup
	make server

start-remote: ## Starts server without db since db is hosted remotely
	make migrateup
	make server

stop: ## Stops every service
	make stopdb
	sleep 2
	sudo systemctl stop docker.socket docker

.PHONY: startdb stopdb createdb dropdb resetdb psql migrateup migratedown migrateup1 migratedown1 sqlc proto test server build start stop start-remote