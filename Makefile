GO := $(shell which go)
GOLINT := $(shell which golint)
GOFMT := $(shell which gofmt)
TARGETS = linux darwin
BUILD_DIR = build
PACKAGES := $(shell find . -maxdepth 1 -type d | grep -v -E 'vendor|\.git|build' | tr -d './')

all: clean prepare vet lint format build

clean:
	@rm -rf $(BUILD_DIR)

prepare:
	@mkdir -p $(BUILD_DIR)

build: $(TARGETS)

vet:
	@echo "Running vet"
	@$(GO) vet ./...

lint:
	@echo "Running golint"
	@$(GOLINT) -set_exit_status $(PACKAGES)

format:
	@echo "Running gofmt"
	@$(GOFMT) -w -l -s .

$(TARGETS):%:
	@GOOS=$@ GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/minesweeper-API-$@ $(COMPILE_FLAGS)
