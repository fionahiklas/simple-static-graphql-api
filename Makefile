
.PHONY: help tidy install_tools lint format generate test test_clean test_report

# from https://suva.sh/posts/well-documented-makefiles/
# sorting from https://stackoverflow.com/questions/14562423/is-there-a-way-to-ignore-header-lines-in-a-unix-sort
# Add double hash comments after every target to provide help text
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | awk 'NR<6{print $0; next}{print $0 | "sort"}'

tidy: ## Ensure modules are downloaded
	go mod tidy

install_tools: tidy ## Install tooling needed by other targets
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install golang.org/x/tools/cmd/goimports
	go install github.com/golang/mock/mockgen

lint: ## Lint codebase
	golangci-lint -E goimports run internal/...

format: ## Format codebase and check imports
	goimports -w internal/

generate: ## Generate mock files
	go generate ./internal/...

clean_test_cache: ## Allows all tests to be forced to run without using cached results
	go clean -testcache

test_clean: clean_test_cache test ## Run tests but clear the cache first

test_report: ## Run unit tests and generate an HTML coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test: ## Run unit tests and check coverage
	@echo "Running unit tests"
	go test ./... -cover