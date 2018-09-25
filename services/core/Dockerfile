FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD core-srv /app/core-srv
ENTRYPOINT [ "/app/core-srv" ]
CMD ["./core-srv"]
