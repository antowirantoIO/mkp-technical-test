# Variables
APP_NAME=golang-clean-architecture
DOCKER_COMPOSE=docker-compose
GO=go
MIGRATE=migrate

# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=golang_clean_architecture
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Default target
.DEFAULT_GOAL := help

## Help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

## Development
dev: ## Run the application in development mode
	$(GO) run cmd/web/main.go

worker: ## Run the worker
	$(GO) run cmd/worker/main.go

build: ## Build the application
	$(GO) build -o bin/$(APP_NAME) cmd/web/main.go

test: ## Run tests
	$(GO) test -v ./test/...

test-coverage: ## Run tests with coverage
	$(GO) test -v -coverprofile=coverage.out ./test/...
	$(GO) tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	rm -rf bin/ coverage.out coverage.html

## Dependencies
deps: ## Download dependencies
	$(GO) mod download

tidy: ## Tidy dependencies
	$(GO) mod tidy

vendor: ## Vendor dependencies
	$(GO) mod vendor

## Database
db-create: ## Create database
	createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

db-drop: ## Drop database
	dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

db-migrate-up: ## Run database migrations up
	$(MIGRATE) -database "$(DB_URL)" -path db/migrations up

db-migrate-down: ## Run database migrations down
	$(MIGRATE) -database "$(DB_URL)" -path db/migrations down

db-migrate-force: ## Force database migration version
	@read -p "Enter migration version: " version; \
	$(MIGRATE) -database "$(DB_URL)" -path db/migrations force $$version

db-migrate-version: ## Show current migration version
	$(MIGRATE) -database "$(DB_URL)" -path db/migrations version

## Docker
docker-build: ## Build Docker image
	docker build -t $(APP_NAME) .

docker-run: ## Run Docker container
	docker run -p 3000:3000 $(APP_NAME)

docker-up: ## Start all services with docker-compose
	$(DOCKER_COMPOSE) up -d

docker-down: ## Stop all services
	$(DOCKER_COMPOSE) down

docker-logs: ## Show logs
	$(DOCKER_COMPOSE) logs -f

docker-restart: ## Restart all services
	$(DOCKER_COMPOSE) restart

docker-rebuild: ## Rebuild and restart services
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) build --no-cache
	$(DOCKER_COMPOSE) up -d

## Swagger
swagger-gen: ## Generate Swagger documentation
	swag init -g cmd/web/main.go -o docs/

swagger-fmt: ## Format Swagger comments
	swag fmt

## Linting and Formatting
fmt: ## Format Go code
	$(GO) fmt ./...

vet: ## Run go vet
	$(GO) vet ./...

lint: ## Run golangci-lint
	golangci-lint run

## Production
prod-build: ## Build for production
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o bin/$(APP_NAME) cmd/web/main.go

prod-deploy: ## Deploy to production (customize as needed)
	@echo "Deploying to production..."
	# Add your deployment commands here

.PHONY: help dev worker build test test-coverage clean deps tidy vendor \
        db-create db-drop db-migrate-up db-migrate-down db-migrate-force db-migrate-version \
        docker-build docker-run docker-up docker-down docker-logs docker-restart docker-rebuild \
        swagger-gen swagger-fmt fmt vet lint prod-build prod-deploy