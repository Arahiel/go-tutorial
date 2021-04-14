FROM golang:1.16.0-alpine3.13
RUN mkdir /websocketApp
ADD . /websocketApp
WORKDIR /websocketApp
RUN go mod download
RUN go build -o main .
CMD ["/websocketApp/main"]