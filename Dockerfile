# Multi stage building strategy for reducing image size.
FROM golang:1.13.5-alpine3.10 AS build-env
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app
WORKDIR /app
# Use "Ultimate Packer for eXecutables". ref: https://upx.github.io/
RUN apk add --no-cache --virtual .go-builddeps git gcc make build-base alpine-sdk upx&& \
    go mod download && \
    go get golang.org/x/lint/golint && \
    make validate && \
    make build-linux && \
    go get github.com/pwaller/goupx && \
    goupx main

# If use TLS connection in container, add ca-certificates following command.
# > RUN apk add --no-cache ca-certificates
FROM alpine:3.10.3
RUN mkdir /app && \
    apk add --no-cache tzdata&& \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
WORKDIR /app
COPY --from=build-env /app/main /app
EXPOSE 80
ENTRYPOINT ["/app/main"]
