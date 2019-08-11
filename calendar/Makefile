GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")

build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build .

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v  ./...

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  golangci-lint run --enable-all

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run cmd/calendar/main.go

generate_proto:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) protoc --go_out=. models/event/event.proto


.PHONY: lint  test build run