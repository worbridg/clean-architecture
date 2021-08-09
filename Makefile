SOURCES := $(CURDIR)/cmd/server
BUILD_DIR := $(CURDIR)/bin
ROUTER_TAG := echo

## default: make syntax vet test
.PHONY: default
default: syntax vet test

## syntax: gofmt
.PHONY: syntax
syntax:
	@test -z "$(shell gofmt -s -l .)" || (gofmt -s -d -l . && exit 1)

## vet: go vet
.PHONY: vet
vet:
	go vet -tags $(ROUTER_TAG) ./...

# test: go test
.PHONY: test
test:
	go test -tags $(ROUTER_TAG) -cover -v ./...

## build: go build
.PHONY: build
build: $(BUILD_DIR)
	go generate ./...
	go build -tags $(ROUTER_TAG) -o $</server $(SOURCES)

$(BUILD_DIR):
	mkdir -p $@

## help: show this messages
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## clean: rm -rf bin
.PHONY: clean
clean:
	-rm -rf $(BUILD_DIR)
