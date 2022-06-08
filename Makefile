# Lint is to fix bad code practices
lint: ## Run linters only.
	@echo "\033[2m→ Running linters...\033[0m"
	@golangci-lint run --config .golangci.yml --fix

TEST_TIMEOUT := 30s
test: ## test: Run unit-tests of the project
	@echo ">  Testing the app..."
	@ go test -short -timeout=$(TEST_TIMEOUT) ./...
	@echo ">  Testing done"
