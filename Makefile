.PHONY: migrate create_migration diff_migration

# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# Construct the database URL
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Target to create a new migration
create_migration:
	@read -p "Enter migration name: " name; \
	atlas migrate new --dir file://database/migrations $$name


# Target to apply migrations
migrate:
	@dotenv -e .env -- atlas schema apply --env gorm -u "$(DB_URL)"

# Target to create a migration based on schema differences
diff_migration:
	@read -p "Enter migration name: " name; \
	dotenv -e .env -- atlas migrate diff --env gorm

# Target to hash migrations
hash_migrations:
	@dotenv -e .env -- atlas migrate hash --dir file://database/migrations

db-seed:
	go run cmd/seeder/main.go

dev:
	air
