# Dockerfile
FROM golang:1.11-alpine3.7


RUN apk --no-cache add tini \
&& apk --no-cache add --virtual build-dependencies git \

# Compile delve

&& git clone https://github.com/derekparker/delve.git $GOPATH/src/github.com/derekparker/delve \
&& cd $GOPATH/src/github.com/derekparker/delve \
&& go run scripts/make.go install \

# Clean up
&& apk del --purge -r build-dependencies \
&& rm -rf $GOPATH/src/* \

ENTRYPOINT [ "dlv" ]
CMD ["dlv", "-h"]


