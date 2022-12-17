include ./.env

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up: migrate-install
	migrate -path ./migrations -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGES_HOST}:${POSTGES_PORT}/${POSTGRES_DB}?sslmode=disable&&query" up

migrate-down: migrate-install
	migrate -path ./migrations -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGES_HOST}:${POSTGES_PORT}/${POSTGRES_DB}?sslmode=disable&&query" down 1