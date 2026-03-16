.PHONY: help lint format test coverage ci clean

.DEFAULT_GOAL := help

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk -F ':.*## ' '{printf "  %-30s %s\n", $$1, $$2}'

lint: ## Run linters
	go tool golangci-lint run
	go vet ./...
	go fix -diff ./...
	gofmt -s -d .

format: ## Format source code
	go fix ./...
	gofmt -s -w .

test: ## Run tests
	go test -race ./...

coverage: ## Run tests with coverage report
	go test -race -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

ci: lint test ## Run CI pipeline (lint, test)

clean: ## Remove generated files
	rm -f coverage.out
