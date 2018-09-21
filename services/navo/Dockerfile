FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD navo-api /app/navo-api
ENTRYPOINT [ "/app/navo-api" ]
CMD ["./navo-api"]
