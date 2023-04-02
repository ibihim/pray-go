.PHONY: build install test run watch-test watch-run clean

BINARY_NAME := prayer
GOBIN ?= $(shell go env GOPATH)/bin

build: clean
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/pray/main.go

install: build
	@echo "Installing $(BINARY_NAME) to $(GOBIN)..."
	@mv $(BINARY_NAME) $(GOBIN)

test:
	@echo "Running tests..."
	@go test -v ./...

run:
	@echo "Running the application..."
	@go run ./cmd/pray/main.go next

watch-test:
	@echo "Watching for file changes to run tests..."
	@find . -name "*.go" | entr -r make test

watch-run:
	@echo "Watching for file changes to run the application..."
	@find . -name "*.go" | entr -r make run

clean:
	@echo "Cleaning up $(BINARY_NAME)..."
	@rm -f $(BINARY_NAME)

