# Multi stage building strategy for reducing image size.
FROM golang:1.14-alpine3.12 AS build-env
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
FROM gcr.io/distroless/base-debian10
COPY --from=build-env /app/main /
EXPOSE 80
ENTRYPOINT ["/main"]
