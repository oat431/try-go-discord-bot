.PHONY: build run test clean tidy

APP_NAME=try-go-discord-bot
MAIN_FILE=cmd/bot/main.go

# Build the application
build:
	@echo "Building try-go-discord-bot..."
	@go build -o bin/$(APP_NAME) $(MAIN_FILE)

# Run the application directly
start:
	@echo "Starting try-go-discord-bot..."
	@go run $(MAIN_FILE)

# Run tests
test:
	@echo "Testing try-go-discord-bot..."
	@go test -v ./...

# Clean the build directory
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf bin/

# Tidy module dependencies
tidy:
	@echo "Tidying..."
	@go mod tidy