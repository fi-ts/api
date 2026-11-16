MAKEFLAGS += --no-print-directory
BUF_VERSION := 1.75.0
PROTODIR := $(shell pwd)

_buf:
	docker run --rm \
		--entrypoint sh \
		-v $(PWD)/..:/workspace \
		-w /workspace/proto \
		bufbuild/buf:$(BUF_VERSION) \
		-c "buf $(CMD) && chown -R $(shell id -u):$(shell id -g) /workspace"

.PHONY: protolint
protolint:
	@$(MAKE)	_buf	CMD="format -w"
	@$(MAKE)	_buf	CMD="lint -v"
	# @$(MAKE)	_buf	CMD="breaking --against 'https://github.com/fi-ts/api.git#branch=main,subdir=proto'"

.PHONY: protoc
protoc: _install-protoc-gen-mcp
	@$(MAKE)	_buf	CMD="generate -v"

.ONESHELL:
.PHONY: _install-protoc-gen-mcp
_install-protoc-gen-mcp:
	cd /tmp
	rm -rf protoc-gen-go-mcp
	# TODO once https://github.com/redpanda-data/protoc-gen-go-mcp/pull/51 is merged, we can use upstream
	# Check out the support-connect-simple branch of https://github.com/majst01/protoc-gen-go-mcp.git
	git clone https://github.com/majst01/protoc-gen-go-mcp.git
	git -C protoc-gen-go-mcp checkout 3f67e6f557ed72f0bbf872861299f577790c6cb0
	cd protoc-gen-go-mcp
	# TODO GOEXPERIMENT can be removed once upgraded to go-1.27
	CGO_ENABLED=0 GOEXPERIMENT=nojsonv2 go build ./cmd/protoc-gen-go-mcp
	cp protoc-gen-go-mcp $(PROTODIR)/
