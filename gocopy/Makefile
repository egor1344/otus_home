GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

build:
	@echo "Building $(GOFILES) files in ./cmd/"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .
	
test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

.PHONY: build get install run test