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

lint: ## run linter
	golangci-lint run --timeout 5m0s ./app/...

gen-oapi-server:
	oapi-codegen --config=oapi-codegen/oapi-config.yml -o backend/oapi/oapi.gen.go openapi/openapi.yml

gen-oapi-client:
	docker compose run --rm openapi-generator-cli generate -g typescript-axios -i openapi/openapi.yml -o ./frontend/src/api/axios --additional-properties=useSingleRequestParameter=true