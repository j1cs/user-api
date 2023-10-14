include .env
export

.PHONY: git-hooks
git-hooks:
	@chmod +x ./.githooks/* && cp -f ./.githooks/* ./.git/hooks/

.PHONY: openapi
openapi: git-hooks
	@./scripts/openapi.sh api internal/v1/port port

.PHONY: mock
mock:
	@./scripts/mock.sh

.PHONY: dev
dev: git-hooks srv
	@./scripts/run.sh "dev"

.PHONY: debug
debug: git-hooks srv
	@./scripts/run.sh "debug"

.PHONY: srv
srv:
	@docker compose -f docker-compose.srv.yml up -d --remove-orphans

.PHONY: run
run:
	@docker compose -f docker-compose.run.yml up --remove-orphans

.PHONY: build
build:
	@docker compose -f docker-compose.run.yml build

.PHONY: down
down:
	@@docker compose -f docker-compose.srv.yml -f docker-compose.run.yml down

.PHONY: pre-coverage
pre-coverage:
	@./scripts/coverage/precoverage.sh

.PHONY: ignore-coverage
ignore-coverage:
	@./scripts/coverage/exclude.sh

.PHONY: coverage
coverage: pre-coverage ignore-coverage
	@go tool cover -func coverage.out

.PHONY: format
format:
	@@./scripts/format.sh

.PHONY: lint
lint:
	@@./scripts/lint.sh