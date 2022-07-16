FROM golang:1.19rc2-alpine3.16
RUN apk update && apk add git alpine-sdk build-base && mkdir /go/src/app
WORKDIR /go/src/app
ADD app/ /go/src/app
