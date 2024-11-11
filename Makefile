build:
	docker-compose build

run:
	docker-compose up

create_migration:
	@go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations postgres "user=escoutdoor password=ivan dbname=bookit sslmode=disable" create $(name) sql

migrations_up:
	@go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations postgres "host=localhost port=3900 user=escoutdoor password=ivan dbname=bookit sslmode=disable" up

migrations_reset:
	@go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations postgres "host=localhost port=3900 user=escoutdoor password=ivan dbname=bookit sslmode=disable" reset
