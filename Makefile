.PHONY: help build build-all build-macos-amd64 build-macos-aarch64 build-linux-amd64 build-linux-aarch64 install test run watch-test watch-run clean

BINARY_NAME := prayer
GOBIN ?= $(shell go env GOPATH)/bin
MAIN_PATH := ./cmd/pray/main.go

help:
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@echo "  help                Show this help message"
	@echo "  build               Build the $(BINARY_NAME) binary for current platform"
	@echo "  build-all           Build the $(BINARY_NAME) binary for all platforms"
	@echo "  build-macos-amd64   Build the $(BINARY_NAME) binary for macOS amd64"
	@echo "  build-macos-aarch64 Build the $(BINARY_NAME) binary for macOS aarch64"
	@echo "  build-linux-amd64   Build the $(BINARY_NAME) binary for Linux amd64"
	@echo "  build-linux-aarch64 Build the $(BINARY_NAME) binary for Linux aarch64"
	@echo "  install             Install the $(BINARY_NAME) binary to $(GOBIN)"
	@echo "  test                Run tests"
	@echo "  run                 Run the application"
	@echo "  watch-test          Watch for file changes and run tests"
	@echo "  watch-run           Watch for file changes and run the application"
	@echo "  clean               Clean up the $(BINARY_NAME) binary"

build: clean
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)

build-all: build-macos-amd64 build-macos-aarch64 build-linux-amd64 build-linux-aarch64

build-macos-amd64:
	@echo "Building $(BINARY_NAME) for macOS amd64..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-macos-amd64 $(MAIN_PATH)

build-macos-aarch64:
	@echo "Building $(BINARY_NAME) for macOS aarch64..."
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-macos-aarch64 $(MAIN_PATH)

build-linux-amd64:
	@echo "Building $(BINARY_NAME) for Linux amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

build-linux-aarch64:
	@echo "Building $(BINARY_NAME) for Linux aarch64..."
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-linux-aarch64 $(MAIN_PATH)

install: build
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

