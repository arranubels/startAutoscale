GOPATH=$(shell pwd)/gopath/

all: build

run:
	GOPATH=${GOPATH} go run *.go

build:
	GOPATH=${GOPATH} go build

get:
	GOPATH=${GOPATH} go get -u      "github.com/aws/aws-sdk-go"
