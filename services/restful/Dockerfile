# FROM alpine:3.2
FROM ironkit/golang:latest

ADD html /html
ADD swigger /swigger
ADD resource /resource
ADD restful-web /restful-web

WORKDIR /

ENTRYPOINT [ "/restful-web" ]

CMD [ "dlv", "--headless", "--listen=:2345", "--api-version=2" ]
# CMD ["dlv", "--headless", "--listen=:2345", "--api-version=2", "exec", "/restful-web"]
# CMD [""]
# ENTRYPOINT ["dlv --listen=:2345 --headless=true --api-version=2 exec ./restful-web" ]


