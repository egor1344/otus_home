build:
	@GOPATH=$(GOPATH) go build .

test:
	@GOPATH=$(GOPATH) go test -v  ./...

lint:
	@GOPATH=$(GOPATH)  golangci-lint run --enable-all

run_server:
	@GOPATH=$(GOPATH) go run server/cmd/main.go

run_client:
	@GOPATH=$(GOPATH) go run client/cmd/main.go --addres=127.0.0.1 --port=3000 --timeout=10


.PHONY: lint  test build run_server, run_server