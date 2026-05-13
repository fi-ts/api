
SHA := $(shell git rev-parse --short=8 HEAD)
GITVERSION := $(shell git describe --long --all)
BUILDDATE := $(shell date -Iseconds)
VERSION := $(or ${VERSION},$(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD))
LOCALBIN ?= $(shell pwd)/bin

all: proto generate test build

release: proto generate test build

.PHONY: proto
proto: protolint
	$(MAKE) -C go clean
	$(MAKE) -C python clean
	$(MAKE) -C proto protoc
	$(MAKE) -C python generate

.PHONY: protolint
protolint:
	$(MAKE) -C proto protolint

.PHONY: generate
generate:
	$(MAKE) -C generate generate

.PHONY: test
test:
	$(MAKE) -C go test

.PHONY: build
build:
	make -C js build VERSION=$(VERSION)
