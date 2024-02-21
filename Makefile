.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  deps      to install dependencies"
	@echo "  dev       to start development"
	@echo "  build     to build binary. Use build_args to set build args. Default is '-o ./release/'"

.PHONY: dev
dev:
	@echo "Starting development service..."
	@go run cmd/main.go

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf ./release ./dist ./db ./logs ./upload ./cmd/db ./cmd/upload ./cmd/logs
	@echo "Cleaned up."

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod verify
	@go mod tidy
	@echo "Dependencies installed."

.PHONY: build
build:
	@echo "Building binary..."
	@bash ./script/build.sh $(build_args)
	@echo "Binary built successfully."
