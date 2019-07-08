GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

build:
	@echo "Building $(GOFILES) to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(GONAME)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) chmod 777 bin/$(GONAME)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/gocopy/
	
test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) golangci-lint run --enable-all

.PHONY: build get install run test