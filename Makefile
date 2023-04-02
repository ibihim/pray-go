.PHONY: help build install test run watch-test watch-run clean

BINARY_NAME := prayer
GOBIN ?= $(shell go env GOPATH)/bin
MAIN_PATH := ./cmd/pray/main.go

help:
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@echo "  help        Show this help message"
	@echo "  build       Build the $(BINARY_NAME) binary"
	@echo "  install     Install the $(BINARY_NAME) binary to $(GOBIN)"
	@echo "  test        Run tests"
	@echo "  run         Run the application"
	@echo "  watch-test  Watch for file changes and run tests"
	@echo "  watch-run   Watch for file changes and run the application"
	@echo "  clean       Clean up the $(BINARY_NAME) binary"

build: clean
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)

install: build clean
	@echo "Installing $(BINARY_NAME) to $(GOBIN)..."
	@mv $(BINARY_NAME) $(GOBIN)

test:
	@echo "Running tests..."
	@go test -v ./...

run:
	@echo "Running the application..."
	@go run $(MAIN_PATH) next

watch-test:
	@if command -v entr > /dev/null; then \
		echo "Watching for file changes to run tests..."; \
		find . -name "*.go" | entr -r make test; \
	else \
		echo "Error: 'entr' command not found. Please install 'entr' to use this feature."; \
	fi

watch-run:
	@if command -v entr > /dev/null; then \
		echo "Watching for file changes to run the application..."; \
		find . -name "*.go" | entr -r make run; \
	else \
		echo "Error: 'entr' command not found. Please install 'entr' to use this feature."; \
	fi

clean:
	@echo "Cleaning up $(BINARY_NAME)..."
	@rm -f $(BINARY_NAME)

