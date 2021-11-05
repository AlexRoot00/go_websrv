.PHONY: build
build:
	go build -v ./cmd/websrv
.DEFAULT_GOAL := build