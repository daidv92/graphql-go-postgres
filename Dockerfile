FROM golang:1.14

RUN go version

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

EXPOSE 8080
EXPOSE 3004