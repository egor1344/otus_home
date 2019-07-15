GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

build:
	@echo "Building $(GOFILES) files in cmd"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) pwd
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build cmd/main.go

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -test.v

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

test_programm:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) ./cmd/main ./test_env env


.PHONY: lint  test build test_programm