.PHONY:

.PHONY: build
build: test ## Build the ports service
	go build -o bin/ports ./internal/port

.PHONY: run
run: ## Run the ports service
	go run ./internal/port/main.go

.PHONY: test
test: docker-up ## Run the tests
	source .env && go test -v ./internal/...

.PHONY: fmt
fmt: ## Run gofmt on all go files
	gofmt -l -w internal/

.PHONY: docker-build
docker-build: ## Build docker image
	docker-compose build

.PHONY: docker-up
docker-up: ## Run docker-compose up
	docker-compose up -d

.PHONY: docker-down
docker-down: ## Run docker-compose down
	docker-compose down

.PHONY: help
help: ## Show make targets
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'