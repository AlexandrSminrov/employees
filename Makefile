include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build
build: ## Docker run
	@docker-compose up -d

run: ## run main.go
	@go run main.go

cover:
	go test -coverprofile cover.out&&go tool cover -html=cover.out

swagger:
	swag init

