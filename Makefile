PACKAGES := $(shell go list ./... | grep -v /vendor/)


.PHONY: run
run:
	swag init -g cmd/main.go
	go run cmd/main.go 

.PHONY: build
build:
	go mod download
	go mod verify
	@echo "Build done."

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)
	@echo "Linters done."