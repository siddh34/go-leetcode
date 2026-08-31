default:
    just

lint:
    go tool golangci-lint run

migrate-lint:
    go tool golangci-lint migrate
