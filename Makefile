app-api=cpf-cnpj-api
compose = docker compose

.PHONY: build 
build:
	$(compose) build

.PHONY: up
up:
	$(compose) up

.PHONY: clean
clean:
	$(compose) down --remove-orphans --volumes

.PHONY: logs
logs:
	$(compose) logs -f --tail=100

lint_version_golang ?= v1.61.0
.PHONY: lint-api
lint-api:
	docker run --rm \
		-w /app \
		-v $(PWD)/$(app-api):/app \
		golangci/golangci-lint:$(lint_version_golang) \
		golangci-lint run -c tools/.golang-ci.yaml --timeout 3m

.PHONY: test
test: build
	$(compose) run --rm api go test ./tests/unit/...
