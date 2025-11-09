.PHONY: run build test clean help dev

# Default target
help:
	@echo "Collaborative Todo List API - Makefile"
	@echo ""
	@echo "Available commands:"
	@echo "  make dev       - Run with hot reload (like npm run dev)"
	@echo "  make run       - Run the application"
	@echo "  make build     - Build the application"
	@echo "  make test      - Run the test script (requires server to be running)"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make deps      - Install dependencies"
	@echo "  make help      - Show this help message"

# Run with hot reload (like npm run dev)
dev:
	@echo "Starting Todo API with hot reload..."
	@echo "Server will auto-restart on code changes"
	@which air > /dev/null || (echo "Air not found. Installing..." && go install github.com/air-verse/air@latest)
	@$(shell go env GOPATH)/bin/air || air

# Run the application
run:
	@echo "Starting Todo API server..."
	go run cmd/api/main.go

# Build the application
build:
	@echo "Building Todo API..."
	go build -o todo-api cmd/api/main.go
	@echo "Build complete: ./todo-api"

# Run tests (requires server to be running)
test:
	@echo "Running API tests..."
	@echo "Make sure the server is running on port 8080"
	@./test_api.sh

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f todo-api
	rm -rf tmp/
	rm -f build-errors.log
	@echo "Clean complete"

# Run with custom port
run-port:
	@echo "Starting Todo API server on port $(PORT)..."
	PORT=$(PORT) go run cmd/api/main.go
