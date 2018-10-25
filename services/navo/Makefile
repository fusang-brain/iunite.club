
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:

	GOOS=linux GOARCH=amd64 go build -o navo-web main.go plugin.go
	go build -o navo-web-local main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t navo-web:latest
