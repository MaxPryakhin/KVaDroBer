SERVER_APP_NAME=kvadrober-server
CLI_APP_NAME=kvadrober-cli

build-server:
	go build -o ${SERVER_APP_NAME} cmd/server/main.go

build-cli:
	go build -o ${CLI_APP_NAME} cmd/cli/main.go

install-deps:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

run-unit-test:
	go test -v ./internal/...

run_e2e_test: build-server
	go test ./test/...

run-test-coverage:
	go test ./... -coverprofile=coverage.out

mock:
	mockgen -source=internal/database/storage/storage.go -destination=internal/database/storage/storage_mock.go -package=storage
	mockgen -source=internal/database/database.go -destination=internal/database/database_mock.go -package=database