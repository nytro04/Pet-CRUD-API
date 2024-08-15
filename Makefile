include .env

build:
	@go build -o bin/api

run: build
	@./bin/api

start-db:
	docker run --name ${POSTGRES_DB_NAME} -e POSTGRES_PASSWORD=${POSTGRES_DB_PASSWORD} -p ${POSTGRES_DB_PORT}:${POSTGRES_DB_PORT} -d postgres
