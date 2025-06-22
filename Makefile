.PHONY: build test lint clean install help release release-snapshot release-check docs docs-serve docs-build

# Build variables
BINARY_NAME=cli-template
BUILD_DIR=bin
MAIN_PATH=./main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

test: ## Run tests
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run golangci-lint
	golangci-lint run

lint-fix: ## Run golangci-lint with auto-fix
	golangci-lint run --fix

clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf dist
	rm -f coverage.out coverage.html
	rm -rf site

deps: ## Download dependencies
	$(GOMOD) download
	$(GOMOD) tidy

install: build ## Install the binary
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

run: ## Run the application
	$(GOCMD) run $(MAIN_PATH)

dev: ## Run in development mode with auto-rebuild
	@echo "Running in development mode..."
	$(GOCMD) run $(MAIN_PATH)

release-check: ## Check GoReleaser configuration
	goreleaser check

release-snapshot: ## Create a snapshot release without publishing
	goreleaser release --snapshot --clean

release: ## Create a release (requires tag)
	goreleaser release --clean

docs-serve: ## Serve documentation locally using Docker
	@command -v docker >/dev/null 2>&1 || { echo "Docker not found. Install Docker Desktop from https://docker.com/products/docker-desktop"; exit 1; }
	docker run --rm -p 8000:8000 -v $(PWD):/docs squidfunk/mkdocs-material

docs-build: ## Build documentation using Docker
	@command -v docker >/dev/null 2>&1 || { echo "Docker not found. Install Docker Desktop from https://docker.com/products/docker-desktop"; exit 1; }
	docker run --rm -v $(PWD):/docs squidfunk/mkdocs-material build --clean --strict

docs: docs-build ## Build and validate documentation

all: clean deps lint test build ## Run all checks and build