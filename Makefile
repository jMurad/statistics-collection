.PHONY: build
build:
	go build -v ./cmd/statserver

.DEFAULT_GOAL := build
