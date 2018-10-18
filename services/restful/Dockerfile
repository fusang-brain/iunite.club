# FROM alpine:3.2
FROM ironkit/golang:latest

ADD html /html
ADD swigger /swigger
ADD restful-web /restful-web
WORKDIR /
ENTRYPOINT [ "/restful-web" ]
