.PHONY: build
build:
	go build -v ./cmd/apiserver
.PHONY: test
test:
	go test -v -race -timeout 30ms ./...
.DEFAULT_GOAL := build
