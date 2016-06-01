clean: ## Clean the project
	rm -rf ./build
	mkdir ./build

test: ## Run tests
	go test -v ./...

coverage: ## Generate test coverage
	./bin/coverage.sh

view-coverage: ## Open coverage report in browser
	go tool cover -html=".coverage-reports/cover.out"

run: ## Build and run the project
	mkdir -p ./build
	go build -o ./build/terragen && ./build/terragen
