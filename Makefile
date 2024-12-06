run-local: build
	@./bin/app -env local
run-dev: build
	@./bin/app -env dev
build:
	@go build -o bin/app ./cmd/app/main.go