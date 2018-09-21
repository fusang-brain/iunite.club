# FROM alpine:3.2
# ADD message-srv /message-srv
# ENTRYPOINT [ "/message-srv" ]
FROM alpine:3.2

RUN mkdir /app

WORKDIR /app

ADD message-srv /app/message-srv

ENTRYPOINT [ "/app/message-srv" ]

CMD ["./message-srv"]
