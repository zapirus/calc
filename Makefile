.PHONY: build
build:
	go build -v ./cmd/calc
.DEFAULT_GOAL := build