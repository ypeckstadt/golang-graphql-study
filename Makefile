SHELL := /bin/bash

#### BUILD
.PHONY: build-server build-client
build-server:
	@echo "building server ..."
	@cd cmd/server && go build
	@echo "build server finished"

build-client:
	@echo "building client ..."
	@cd cmd/client && go build
	@echo "build client finished"

#### RUN
.PHONY: run-server run-client
run-server: build-server
	@echo "Starting server ..."
	@cd env && source export-all.sh env.dev && cd ../cmd/server && ./server

run-client: build-client
	@echo "Starting server ..."
	@cd env && source export-all.sh env.dev && cd ../cmd/client && ./client

##### ENVIRONMENT
.PHONY: env-up env-up env-reset
env-up:
	@echo "Starting environment"
	@cd fixtures && docker-compose -f docker-compose.yml up --force-recreate --build -d
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose -f docker-compose.yml down -v
	@echo "Environment down"

env-reset: env-down env-up

##### GRAPHQL
.PHONY: ql-gen
ql-gen:
	@go run github.com/99designs/gqlgen generate --verbose

##### SQLBOILER - DATABASE
.PHONY: boiler-update
boiler-update:
	@echo "Removing models"
	@rm -rf models
	@echo "Generating models ..."
	@sqlboiler psql --struct-tag-casing camel -o ./pkg/boiled -p boiled
	@echo "SQLBOILER models updated"
