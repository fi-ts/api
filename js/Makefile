VERSION_STRIPPED := $(patsubst v%,%,$(VERSION))

.PHONY: install
install:
	bun upgrade || true
	bun install -dd

.PHONY: clean
clean:
	rm -rf fits

.PHONY: build
build: install
ifeq ($(CI),true)
	cd .. &&\
	yq e -i -o=json ".version=\"$(VERSION_STRIPPED)\"" package.json &&\
	yq e -o=json ".version" package.json
endif
	cd .. && bun run build
