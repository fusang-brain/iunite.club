# FROM alpine:3.2
# ADD user-srv /user-srv
# ENTRYPOINT [ "/user-srv" ]

FROM alpine:3.2

RUN mkdir /app

WORKDIR /app

ADD user-srv /app/user-srv

ENTRYPOINT [ "/app/user-srv" ]

CMD ["./user-srv"]
