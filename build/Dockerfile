FROM golang:1.13.1-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base
RUN set -x \
    # go get mux and mongo
    && go get github.com/gorilla/mux \
    && go get go.mongodb.org/mongo-driver/mongo 
WORKDIR /go/src