
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	ironic --path="./"

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-srv .

test:
	go test -v ./... -cover

#docker:
#	docker build . -t user-srv:latest

run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns \
	-e DB_HOST="127.0.0.1:27018" \
	-e DB_NAME="unite" \
	user-srv
