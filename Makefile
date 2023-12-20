build: ## Build docker compose
	docker compose build

rebuild: ## Run docker compose with --no-cache
	docker compose build --no-cache;

up: ## Run docker compose
	docker compose up;

up-d: ## Run docker compose -d
	docker compose up -d;

down: ## Run docker compose down
	docker compose down;