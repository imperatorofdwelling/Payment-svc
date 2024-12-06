run-local: build
	@./bin/app -env local
build:
	@go build -o bin/app ./cmd/app/main.go