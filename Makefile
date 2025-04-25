# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=url-shortener

# Run the application
run:
	$(GORUN) cmd/server/main.go

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Install dependencies
install:
	$(GOCMD) mod tidy
