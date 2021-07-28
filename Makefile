PROJECTNAME=$(shell basename $(PWD))
GO_LDFLAGS=-ldflags "-X main.Version=`git tag --sort=-version:refname | head -n 1`"
PLATFORMS=darwin windows linux

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
	find . -type f -name '*.go' -not -path '*/\.*' -exec sed -i '' "s/go-create-binary-template/${PROJECTNAME}/g" {} +
	sed -i '' "s/go-create-binary-template/${PROJECTNAME}/g" go.mod

### build: To build the binary in bin directory
.PHONY: build
build:
	$(foreach GOOS, $(PLATFORMS),\
	$(shell mkdir -p bin && cd bin && env GOOS=$(GOOS) GOARCH=amd64 go build $(GO_LDFLAGS) -o ${PROJECTNAME}-${GOOS} ../ ))
	mv bin/${PROJECTNAME}-{windows,windows.exe}
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
