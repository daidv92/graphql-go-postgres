FROM golang:1.14

RUN go version

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get app
RUN go install
ENTRYPOINT ["/go/bin/app"]