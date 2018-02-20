FROM golang

RUN mkdir /go/src/public-order-bot
ADD . /go/src/public-order-bot
WORKDIR /go/src/public-order-bot

RUN go get -u github.com/golang/dep/...
RUN dep ensure
