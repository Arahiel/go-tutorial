FROM golang:1.16.0-alpine3.13
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8081
RUN go build -o main ./server/
CMD ["/app/main"]