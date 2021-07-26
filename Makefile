PROJECTNAME=$(shell basename $(PWD))
COMMIT=$(shell git rev-parse --short HEAD)
GO_LDFLAGS=-ldflags "-X main.Version=v1.0.0|commit="$(COMMIT)""
PLATFORMS=darwin windows

# Make is verbose in Linux. Made it silent
MAKEFLAGS += --silent

## usage: make [option]
## Options:
.PHONY : help
help : Makefile
	echo "Choose a command run in $(PROJECTNAME):"
	sed -n 's/^## //p' $<
	sed -n 's/^###/ /p' $<

### bootstrap: To bootstrap the template
.PHONY: bootstrap
bootstrap:
	find . -type f -not -path '*/\.*' -exec sed -i "s/go-create-binary-template/`basename $PWD`/g" {} +

### build: To build the binary in bin directory
.PHONY: build
build:
	$(foreach GOOS, $(PLATFORMS),\
	$(shell export GOOS=$(GOOS); mkdir -p bin && cd bin && go build $(GO_LDFLAGS) ../ ))
	echo "Binaries are successfully built."

### clean: To remove the binaries 
.PHONY: clean
clean:
	rm -rfv bin/

### sanity-check: To execute lint, errors and security checks
.PHONY: sanity-check
sanity-check:
	go vet 
	gosec ./...
