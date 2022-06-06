.PHONY: build

build:
	go build -v ./cmd/apiserver

.PHONY: test

test:
	get test -v -race -timeout 30s ./ ...

.DEFAULT_GOAL := build