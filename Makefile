GO := $(shell which go)
GOLINT := $(shell which golint)
GOFMT := $(shell which gofmt)
TARGETS = linux darwin
BUILD_DIR = build
PACKAGES := $(shell find . -maxdepth 1 -type d | grep -v -E 'vendor|\.git|build' | tr -d './')
COVERED_PKGS = github.com/marcelog/minesweeper-API/endpoints,github.com/marcelog/minesweeper-API/server,github.com/marcelog/minesweeper-API/user,github.com/marcelog/minesweeper-API/state,github.com/marcelog/minesweeper-API/game

all: clean prepare vet lint format test build show_coverage

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

test:
	@$(GO) test ./... -v -covermode=set -coverpkg=$(COVERED_PKGS) -coverprofile=coverage.out

show_coverage:
	@$(GO) tool cover -html=coverage.out

$(TARGETS):%:
	@GOOS=$@ GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/minesweeper-API-$@ $(COMPILE_FLAGS)
