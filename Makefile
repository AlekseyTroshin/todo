build:
	@go build -o bin/todo cmd/todo/main.go

run: build
	@./bin/todo