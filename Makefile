build:
    go build -o bin/main ./cmd/app

run:
    go run ./cmd/app

test:
    go test -v -cover ./...

lint:
    golangci-lint run
