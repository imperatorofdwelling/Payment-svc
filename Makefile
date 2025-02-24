include .env

run-local: build
	@./bin/app -env local
run-dev: build
	@./bin/app -env dev
build:
	@go build -o bin/app ./cmd/app/main.go

### Docker ###
docker-local: yml-convert-local
	@docker compose --env-file .env -f ./docker/local/docker-compose.yml -p iod-payment up --build -d
	@$(MAKE) migrate-up-docker-local
docker-dev: yml-convert-dev
	@docker compose --env-file .env -f ./docker/dev/docker-compose.yml -p iod-payment up --build -d
	@$(MAKE) migrate-up-docker-dev
docker-prod: yml-convert-prod
	@docker compose --env-file .env -f ./docker/prod/docker-compose.yml -p iod-payment up --build -d

### Migrate ###
#migrate-up: create-env
#	@go run cmd/migrator/main.go up
#migrate-down:
#	@go run cmd/migrator/main.go down
#migration-create:
#	@migrate create -ext sql -dir .\internal\storage\postgres\migrations -seq $(filter-out $@,$(MAKECMDGOALS))
migrate-down-docker-local:
	@docker compose --env-file ./.env -f ./docker/local/docker-compose.yml -p iod-payment run --rm migrate-down-payment
migrate-up-docker-local:
	@docker compose --env-file ./.env -f ./docker/local/docker-compose.yml -p iod-payment run --rm migrate-payment
migrate-down-docker-dev:
	@docker compose --env-file ./.env -f ./docker/dev/docker-compose.yml -p iod-payment run --rm migrate-down-payment
migrate-up-docker-dev:
	@docker compose --env-file ./.env -f ./docker/dev/docker-compose.yml -p iod-payment run --rm migrate-payment

### YML config convert to .env (for docker compose usage) ###
create-env:
	@go build -o bin/ymlConverter ./cmd/ymlConverter/main.go

yml-convert-local: create-env
	@./bin/ymlConverter -env local
yml-convert-dev: create-env
	@./bin/ymlConverter -env dev
yml-convert-prod: create-env
	@./bin/ymlConverter -env prod