# Makefile for Go + Vite Docker project

.PHONY: help dev prod build-dev build-prod up-dev up-prod down logs clean restart-dev restart-prod

# Default target
help:
	@echo "Available commands:"
	@echo "  dev          - Start development environment (detached)"
	@echo "  prod         - Start production environment (detached)"
	@echo "  build-dev    - Build development images"
	@echo "  build-prod   - Build production images"
	@echo "  up-dev       - Start dev environment in foreground"
	@echo "  up-prod      - Start prod environment in foreground"
	@echo "  down         - Stop all services"
	@echo "  logs         - View logs from all services"
	@echo "  logs-backend - View backend logs only"
	@echo "  logs-frontend- View frontend logs only"
	@echo "  clean        - Remove all containers, volumes, and images"
	@echo "  restart-dev  - Restart development environment"
	@echo "  restart-prod - Restart production environment"

# Development commands
dev:
	docker compose --profile dev up -d

build-dev:
	docker compose --profile dev build

up-dev:
	docker compose --profile dev up

restart-dev:
	docker compose --profile dev restart

# Production commands
prod:
	docker compose --profile prod up -d

build-prod:
	docker compose --profile prod build -d

up-prod:
	docker compose --profile prod up

restart-prod:
	docker compose --profile prod restart

# General commands
down:
	docker compose down

logs:
	docker compose logs -f

logs-backend:
	docker compose logs -f backend-dev backend

logs-frontend:
	docker compose logs -f frontend-dev frontend

# Clean up everything
clean:
	docker compose down -v --remove-orphans
	docker system prune -af --volumes

# Rebuild and start development
rebuild-dev:
	docker compose --profile dev down
	docker compose --profile dev build --no-cache
	docker compose --profile dev up -d

# Rebuild and start production
rebuild-prod:
	docker compose --profile prod down
	docker compose --profile prod build --no-cache
	docker compose --profile prod up -d