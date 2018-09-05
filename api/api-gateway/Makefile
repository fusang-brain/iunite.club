
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	ironic --path="./"

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t api-gateway .

test:
	go test -v ./... -cover

#docker:
#	docker build . -t user-srv:latest

run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns \
	api-gateway
