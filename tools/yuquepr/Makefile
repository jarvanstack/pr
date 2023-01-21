# config
bindir = ./bin
releaseDir = ./release
mainFile = ./main.go

# var
make_dir:=$(shell pwd)
app_name:=$(shell basename $(make_dir))
version=$(shell git describe --tags --always)

## build: Builds the project
.PHONY: build
build: 
	mkdir -p $(bindir)
	GOOS=linux GOARCH=amd64 go build -o $(bindir)/$(app_name).linux $(mainFile)

## install: Installs the project to $GOPATH/bin/$(app_name)
.PHONY: install
install: clean build
	cp $(bindir)/$(app_name).linux $(GOPATH)/bin/$(app_name)
	cp ypr $(GOPATH)/bin/ypr
## clean: Cleans the project
.PHONY: clean
clean:
	rm -rf $(bindir)/*
	rm -rf $(releaseDir)/*
	rm -rf images/*
	go clean

## help: Show this help info.
.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS