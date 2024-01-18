PROJECT_NAME := Template
EXE_NAME := template

.PHONY: test ## run tests
test:
	@echo "Running tests..."
	@go clean -testcache
	@go test -v -coverprofile=coverage ./...
	@echo "Creating coverage report..."
	@go tool cover -html=coverage -o coverage.html
	@echo "Writing coverage report to coverage.html..."
	@echo "Done."