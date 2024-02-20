.PHONY: build

BUILD_DIR = bin

build:
	make prerequisites
	go build -o bin -v ./cmd/api

prerequisites:
	if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)

run:
	make build
	make swagger
	$(BUILD_DIR)\api.exe

swagger:
	swag init -g internal/handlers/server.go -o docs

.DEFAULT_GOAL := build