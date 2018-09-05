FROM alpine:3.2

RUN mkdir /app

WORKDIR /app

ADD user-srv /app/user-srv

# ENTRYPOINT [ "/user-srv" ]
CMD ["./user-srv"]
