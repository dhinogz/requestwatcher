include .env

PSQL_DSN = "postgres://${PSQL_USER}:${PSQL_PASSWORD}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_NAME}?sslmode=${PSQL_SSLMODE}"

build:
	@echo "Building Go application..."
	@templ generate
	@sqlc generate
	@go build -o ./bin/web ./cmd/web

run:
	./bin/web


test:
	@echo "Testing..."
	@go test ./... -v

dev: templ/gen db/up db/migrate
	air

db/up:
	@docker compose up -d
	@echo "Starting DB..."
	@sleep 1

db/migrate:
	@echo "Migrating SQL schema..."
	@goose -dir="./sql/schemas" postgres $(PSQL_DSN) up

db/gen:
	@sqlc generate

db/psql:
	psql $(PSQL_DSN)

db/down:
	@docker compose down -v
	@echo "Removing DB..."

templ/gen:
	@templ generate

templ/watch:
	templ generate --watch
