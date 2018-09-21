
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	# protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto
	ironic proto --path "./proto"

build: proto

	GOOS=linux GOARCH=amd64 go build -o navo-api main.go plugin.go
	go build -o navo-api-local main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t navo-api:latest
