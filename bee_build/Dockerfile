FROM golang:1.15.3-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base

ENV GO111MODULE=on
RUN set -x \
    && go get -u github.com/astaxie/beego \
    && go get -u github.com/beego/bee \
    && bee version
WORKDIR /go/src
