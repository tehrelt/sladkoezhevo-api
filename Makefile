.PHONY: build

BUILD_DIR = bin

build:
	make prerequisites
	go build -o bin -v ./cmd/api

prerequisites:
	if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)

run:
	make build
	$(BUILD_DIR)\api.exe

.DEFAULT_GOAL := build