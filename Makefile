install-deps:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

run-unit-test:
	go test -v ./internal/...

run-test-coverage:
	go test ./... -coverprofile=coverage.out