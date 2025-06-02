COMPOSE = docker-compose
BINARY=app
GOPATH=$(HOME)/go
LINT_PATH=$(GOPATH)/bin/golangci-lint
# $HOME/go/bin/golangci-lint --version
# or $GOPATH/bin/golangci-lint --version
# or $(GOPATH)/bin/golangci-lint
up:
	$(COMPOSE) down
	$(COMPOSE) up -d

build:
	go build -o bin/main ./cmd/app

run:
	go run ./cmd/app

test:
	go test -v -cover ./...


lint: $(LINT_TARGET)
	@echo "==> Linting Go code..."
	@$(LINT_PATH) run --config ./configs/.golangci.yml ./internal/... ./cmd/...

installLint:
	@echo "==> Installing golangci-lint..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

installDebugging:
	go install github.com/go-delve/delve/cmd/dlv@latest