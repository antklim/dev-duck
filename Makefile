.PHONY: build
build: go-build ## Build Go services

.PHONY: clean
clean: go-clean ## Clean build cache and dependencies

.PHONY: test
test: go-test ## Run tests

.PHONY: test-bench
test-bench: go-test-bench ## Run benchmark tests

.PHONY: benchstat
benchstat: go-benchstat ## Benchmark statistics

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build:
	go build -o build/ -v ./...

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	@echo "Cleaning build cache..."
	go clean -cache

go-clean-test-cache:
	@echo "Cleaning test cache..."
	go clean -testcache

go-clean-deps:
	@echo "Cleaning dependencies..."
	go mod tidy

go-deps:
	@echo "Installing dependencies..."
	go mod download

go-test:
	@echo "Running tests..."
	go test -v ./...

go-test-bench:
	go test -bench=. -count=10 | tee addAll.txt

go-benchstat:
	~/go/bin/benchstat addAll.txt

.DEFAULT_GOAL := help