
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:

	GOOS=linux GOARCH=amd64 go build -o restful-web -gcflags "all=-N -l" main.go plugin.go
	go build -o restful-web-local main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t restful-web:latest
