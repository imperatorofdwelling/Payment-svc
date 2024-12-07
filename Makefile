run-local: build
	@./bin/app -env local
run-dev: build
	@./bin/app -env dev
build:
	@go build -o bin/app ./cmd/app/main.go

### Docker ###
docker-local: yml-convert-local
	@docker compose --env-file .env -f ./docker/local/docker-compose.yml -p iod-payment up --build -d
docker-dev: yml-convert-dev
	@docker compose --env-file .env -f ./docker/dev/docker-compose.yml -p iod-payment up --build -d
docker-prod: yml-convert-prod
	@docker compose --env-file .env -f ./docker/prod/docker-compose.yml -p iod-payment up --build -d

### YML config convert to .env (for docker compose usage) ###
create-env:
	@go build -o bin/ymlConverter ./cmd/ymlConverter/main.go
yml-convert-local: create-env
	@./bin/ymlConverter -env local
yml-convert-dev: create-env
	@./bin/ymlConverter -env dev
yml-convert-prod: create-env
	@./bin/ymlConverter -env prod