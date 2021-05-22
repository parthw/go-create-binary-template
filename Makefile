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
	mkdir bin