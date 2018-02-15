FROM golang

RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/...
RUN dep ensure