#!/usr/bin/env make
.DEFAULT_GOAL := help
.SHELLFLAGS   := -eou pipefail

.PHONY: build
build: ## Build package.
	@go build .

.PHONY: test
test: test-unit test-integration ## Run all tests.

.PHONY: test-integration
test-integration: ## Run integration tests.
	INTEGRATION_TEST=1 go test --fullpath -run "TestIntegration" ./...

.PHONY: test-unit
test-unit: ## Run unit tests.
	UNIT_TEST=1 go test --fullpath ./...

.PHONY: help
help: ## Show help/usage.
	@grep -E '^[%a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
