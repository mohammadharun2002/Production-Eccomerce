BINARY := bin/api

ifeq ($(OS),Windows_NT)
	BINARY := bin/api.exe
endif

.PHONY: build run test clean tidy

build:
	@go build -o $(BINARY) ./cmd

run: build
	@$(BINARY)

test:
	@go test -v ./...

tidy:
	@go mod tidy

clean:
	@go clean
	@rm -rf bin
