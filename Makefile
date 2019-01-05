GO := $(shell which go)
TARGETS = linux darwin
BUILD_DIR = build

all: clean prepare build

clean:
	rm -rf $(BUILD_DIR)

prepare:
	mkdir -p $(BUILD_DIR)

build: $(TARGETS)

$(TARGETS):%:
	@GOOS=$@ GOARCH=amd64 go build -o $(BUILD_DIR)/minesweeper-API-$@ $(COMPILE_FLAGS)
