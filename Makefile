# Go parameters
GO        := go
GOBUILD   := $(GO) build
GOTEST    := $(GO) test
GOCLEAN   := $(GO) clean
GOGET     := $(GO) get
BINARY    := blockchain
MAIN      := ./cmd/main.go

all: build

.PHONY: build
build:
	@echo "Building..."
	$(GOBUILD) -o $(BINARY) $(MAIN)

.PHONY: run
run: build
	@echo "Running..."
	./$(BINARY)

.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) ./...

.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY)

.PHONY: get
get:
	@echo "Fetching dependencies..."
	$(GOGET) -v ./...
