SHELL := bash
.ONESHELL:
MAKEFLAGS += --no-builtin-rules

export APP_NAME := $(if $(APP_NAME),$(APP_NAME),employees)
export DOCKER_REPOSITORY := $(if $(DOCKER_REPOSITORY),$(DOCKER_REPOSITORY),employees)
#export VERSION := $(if $(VERSION),$(VERSION),$(if $(COMMIT_SHA),$(COMMIT_SHA),$(shell git rev-parse --verify HEAD)))
export DOCKER_BUILDKIT := 1

CURRENT_GIT_BRANCH := $(shell git branch --show-current)
MIGRATE_DSN := "postgres://pqadmin:pass@localhost:5432/employeesPQ?sslmode=disable"
NOCACHE := $(if $(NOCACHE),"--no-cache")

#.PHONY: help
#help: ## List all available targets with help
#	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)
#		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build app
	@echo ${APP_NAME}
	@docker-compose up -d
# psql -h localhost -d userstoreis -U pqadmin -p 5432 -a -q -f /home/jobs/Desktop/resources/postgresql.sql