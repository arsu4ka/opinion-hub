.PHONY: build
build:
	go build -o bin/app ./cmd/app/main.go

.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: start
start:
	go build -o bin/app ./cmd/app/main.go && bin/app

.PHONY: test
test:
	go test -v -race ./...

.DEFAULT_GOAL := build
