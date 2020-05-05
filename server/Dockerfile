FROM golang:1.13-alpine

WORKDIR /backend
COPY . .
ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/cespare/reflex

COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]