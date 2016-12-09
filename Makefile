clean: ## Clean the project
	rm -rf ./build
	mkdir ./build

test: ## Run tests
	@go test ./...

tdd: ## Watch files, running tests whenever they change
	@./bin/tdd.sh "*.go"

coverage: ## Generate test coverage
	@./bin/coverage.sh

view-coverage: ## Open coverage report in browser
	@go tool cover -html=".coverage-reports/cover.out"

run: ## Build and run the project
	@mkdir -p ./build
	go build -o ./build/terragen && . ./conf/common.sh && . ./conf/dev.sh && ./build/terragen

deploy-web-local: ## Deploys the web server so that the service picks it up automatically
	./bin/deploy-web-local.sh
