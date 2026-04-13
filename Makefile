.PHONY: dev dev-backend dev-frontend build test lint clean sqlc-gen api-types

dev:
	@echo "Starting backend and frontend..."
	$(MAKE) dev-backend &
	$(MAKE) dev-frontend &
	wait

dev-backend:
	cd backend && go run ./cmd/api

dev-frontend:
	cd frontend && yarn dev

build:
	cd backend && go build -o bin/api ./cmd/api
	cd frontend && yarn build

test:
	cd backend && go test ./... -v
	cd frontend && yarn lint

lint:
	cd backend && golangci-lint run
	cd frontend && yarn lint

clean:
	rm -rf backend/bin
	rm -rf frontend/.next frontend/out

compose-up:
	docker compose up --build

compose-down:
	docker compose down

sqlc-gen:
	cd backend && sqlc generate

api-types:
	npx openapi-typescript backend/api/openapi.yml -o frontend/src/lib/types.ts
