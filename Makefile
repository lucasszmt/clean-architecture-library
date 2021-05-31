build:
	@echo "Building"
	go run main.go

MIGRATIONS_PATH="$(shell pwd)/database/migrations"
DB_DSN="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

migrations:
	docker run -v $(MIGRATIONS_PATH):/migrations --network host migrate/migrate -path=/migrations -database $(DB_DSN) up

#Creates the migration file already named
create_migration:
	docker run -v $(MIGRATIONS_PATH):/migrations --user $(shell id -u):$(shell id -g) \
					--network host migrate/migrate create -ext sql -dir ./migrations -seq $(MIG_NAME)