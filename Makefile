export:
	@while read LINE; do export "$LINE"; done < go.env

create-migrations:
	@migrate create -ext sql -dir db -seq create_db_initial

migrate:
	@migrate -database ${POSTGRESQL_URL} -path db up

docker:
	@docker-compose up -d

dockerdown:
	@docker-compose down

run:
	@echo "---- Running Application ----"
	@go run .