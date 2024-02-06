.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  dev       to start development"
	@echo "  build     to build binary. Use build_args to set build args. Default is '-o ./release/'"

.PHONY: dev
dev:
	@echo "Starting development service..."
	@go run cmd/main.go

.PHONY: build
build:
	@echo "Building binary..."
	@bash ./script/build.sh $(build_args)
