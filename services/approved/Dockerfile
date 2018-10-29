FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD approved-srv /app/approved-srv
ENTRYPOINT [ "/app/approved-srv" ]
CMD ["./approved-srv"]
