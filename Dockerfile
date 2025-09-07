# Multi stage building strategy for reducing image size.
FROM public.ecr.aws/docker/library/golang:1.23.4 AS builder
ENV GO111MODULE=on \
    GOPATH=/go \
    GOBIN=/go/bin \
    PATH=/go/bin:$PATH

# Set working directory
WORKDIR /app

# Install each dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

# COPY main module
COPY . /app

# Check and Build
RUN make validate && \
    make build-linux

### If use TLS connection in container, add ca-certificates following command.
### > RUN apt-get update && apt-get install -y ca-certificates
FROM public.ecr.aws/debian/debian
COPY --from=builder /app/bin/main /
EXPOSE 80
ENTRYPOINT ["/main"]
