PACKAGES := $(shell go list ./... | grep -v /vendor/)

.PHONY: run
run:
	swag init -g cmd/server/main.go
	go run cmd/server/main.go 

.PHONY: build
build:
	go mod download
	go mod verify
	@echo "Build done."

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)
	@golangci-lint run -E errcheck
	@echo "Linters done."