FROM alpine:3.2
RUN mkdir /app
WORKDIR /app
ADD report-srv /app/report-srv
ENTRYPOINT [ "/app/report-srv" ]
CMD ["./report-srv"]
