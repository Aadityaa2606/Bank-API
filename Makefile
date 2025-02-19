include .env

POSTGRES_CONTAINER=postgres
POSTGRES_PASSWORD ?= $(or $(POSTGRES_PASSWORD),password)
DB_NAME ?= $(or $(DB_NAME),testdb)
DB_USER ?= $(or $(DB_USER),root)
DB_PORT=5432
DB_HOST=localhost
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgresql://$(DB_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
GOOSE_DIR=./db/migration

postgres:
	sudo docker run --name $(POSTGRES_CONTAINER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p $(DB_PORT):$(DB_PORT) -d postgres:17-alpine

createdb:
	sudo docker start $(POSTGRES_CONTAINER) || true
	sudo docker exec -it $(POSTGRES_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	sudo docker start $(POSTGRES_CONTAINER) || true
	sudo docker exec -it $(POSTGRES_CONTAINER) dropdb --username=$(DB_USER) $(DB_NAME)

migrateup:
	goose $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" -dir $(GOOSE_DIR) up

migratedown:
	goose $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" -dir $(GOOSE_DIR) down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc printenv
