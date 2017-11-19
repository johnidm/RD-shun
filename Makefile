.PHONY: build lint run test deps vet

# GOPATH := /Users/Computer/Projects/RD/shun
# export GOPATH

default: run

build: vet
	go build -v -o ./bin/main ./main.go

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

run: 
	go run ./main.go

test:
	go test ./main_test.go

deps:
	rm -dRf ./vendor/src
	GOPATH=${PWD}/vendor go get -d -u -v \
	github.com/gin-gonic/gin \
	
	
vet:
	go vet ./...