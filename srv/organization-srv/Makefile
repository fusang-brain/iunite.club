
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t organization-srv .

test:
	go test -v ./... -cover

run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns \
	-e DB_HOST="127.0.0.1:27017" \
	-e DB_NAME="unite" \
	organization-srv