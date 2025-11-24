.PHONY: dev build clean test

GO_CMD=go
GO_RUN=$(GO_CMD) run
GO_TEST=$(GO_CMD) test
MAIN_FILE=main.go

NPM_CMD=npm
NPM_RUN=$(NPM_CMD) run

dev:
	@echo "Starting development environment..."
	@echo "Starting Go server and Tailwind CSS dev process..."
	@$(MAKE) -j2 dev-go dev-tailwind

dev-go:
	@echo "Starting Go server..."
	@$(GO_RUN) $(MAIN_FILE)

dev-tailwind:
	@echo "Starting Tailwind CSS dev process..."
	@$(NPM_RUN) dev

build:
	@echo "Building Go binary..."
	@$(GO_BUILD) -o app $(MAIN_FILE)

clean:
	@echo "Cleaning up..."
	@rm -f app

test:
	@echo "Running all tests..."
	@$(GO_TEST) ./... -count=1
