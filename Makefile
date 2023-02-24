run:build
	@./bin/go-hello

build:
	@go build -o bin/go-hello

test:
	@go test -v ./...


help:
	@echo "run --run go-hello"
	@echo "build --build go-hello"
	@echo "test --test go-hello"

.PHONY:run build test

