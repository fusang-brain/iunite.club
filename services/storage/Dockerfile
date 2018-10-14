FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD storage-srv /app/storage-srv
ENTRYPOINT [ "/app/storage-srv" ]
CMD ["./storage-srv"]
