.PHONY: build
build:
	go build -o bin -v ./cmd/api

run:
	make build
	.\bin\api.exe

.DEFAULT_GOAL := build