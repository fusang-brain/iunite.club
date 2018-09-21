# FROM alpine:3.2
# ADD organization-srv /organization-srv
# ENTRYPOINT [ "/organization-srv" ]

FROM alpine:3.2

RUN mkdir /app

WORKDIR /app

ADD organization-srv /app/organization-srv

ENTRYPOINT [ "/app/organization-srv" ]

CMD ["./organization-srv"]
