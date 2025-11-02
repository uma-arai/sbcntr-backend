# === builder: 依存関係生成用 ===
FROM public.ecr.aws/docker/library/golang:1.23.4 AS builder
ENV GO111MODULE=on \
    GOPATH=/go \
    GOBIN=/go/bin \
    PATH=/go/bin:$PATH
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
COPY . /app

# ビルドを実行
RUN make validate && \
    make build-linux

# === runner: 本番イメージ ===
### If use TLS connection in container, add ca-certificates following command.
### > RUN apt-get update && apt-get install -y ca-certificates
FROM public.ecr.aws/debian/debian as runner

# ハンズオンで利用する名前解決用にnslookupを導入(本番向けイメージには不要)
RUN apt-get update && apt-get install -y dnsutils && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/main /
EXPOSE 80
ENTRYPOINT ["/main"]
