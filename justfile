# Default recipe - show available commands
default:
    @just --list --unsorted

# run tests
test:
    go test -count=1 ./...

# run tests with verbose output
test-verbose:
    go test -v -count=1 ./...

# run tests with coverage
cover:
    go test -cover -count=1 ./...

# run tests and show per-function coverage
cover-func:
    go test -coverprofile=/tmp/cliout_cover.out -count=1 . && go tool cover -func=/tmp/cliout_cover.out

# open coverage report in browser
cover-html:
    go test -coverprofile=/tmp/cliout_cover.out -count=1 . && go tool cover -html=/tmp/cliout_cover.out

# run linter
lint:
    golangci-lint run ./...

# run a specific example (e.g. just run-example themes)
run-example name:
    go run ./examples/{{ name }}/

# format code
fmt:
    gofmt -w .

# vet code
vet:
    go vet ./...

# run all checks (fmt, vet, lint, test)
check: fmt vet lint test
