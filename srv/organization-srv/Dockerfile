FROM alpine:3.2

RUN mkdir /app

WORKDIR /app

ADD organization-srv /app/organization-srv

# ENTRYPOINT [ "/user-srv" ]
CMD ["./organization-srv"]