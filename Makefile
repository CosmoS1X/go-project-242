run:
	go run cmd/hexlet-path-size/main.go

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

vuln:
	govulncheck ./...

.PHONY: fmt vet build test

.DEFAULT_GOAL := build
