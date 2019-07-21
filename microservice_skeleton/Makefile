GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

build:
	@echo "Building $(GOFILES) files "
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -test.v

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

run_programm:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) ./microservice_skeleton


.PHONY: lint  test build run_programm