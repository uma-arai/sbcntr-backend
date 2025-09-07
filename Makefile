SHELL := /bin/sh

# Goコマンド
GOCMD         = go
# ビルド後の出力先ディレクトリ
BUILD_DIR     = bin
# ビルドするバイナリの名前
BINARY_NAME   = main
# メインのエントリーポイント
MAIN_PACKAGE  = ./main.go

.PHONY: all validate build test test-verbose test-coverage test-coverage-html run clean update-deps build-linux

# allターゲットでは「validate → build → run」を一括実行
all: validate build run

##
# 検証系: fmt, vet, test
##
validate:
	@echo "==> Running go fmt"
	@$(GOCMD) fmt ./...

	@echo "==> Running go vet"
	@$(GOCMD) vet ./...

	@echo "==> Running golangci-lint"
	@golangci-lint run

#   必要に応じてテストを実行する
# 	@echo "==> Running tests"
# 	@$(GOCMD) test -v ./...

##
# ビルド & 実行
##
build:
	@echo "==> Building..."
	@mkdir -p $(BUILD_DIR)
	@$(GOCMD) build -ldflags "-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run:
	@echo "==> Running..."
	@$(BUILD_DIR)/$(BINARY_NAME)

##
# テスト関連
##
test:
	@echo "==> Running tests"
	@$(GOCMD) test ./...

test-verbose:
	@echo "==> Running tests (verbose)"
	@$(GOCMD) test -v ./...

test-coverage:
	@echo "==> Running tests with coverage"
	@$(GOCMD) test -cover ./...

test-coverage-html:
	@echo "==> Running tests with coverage and generating HTML report"
	@$(GOCMD) test -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

##
# 不要ファイルの削除
##
clean:
	@echo "==> Cleaning build outputs"
	@rm -rf $(BUILD_DIR)

##
# 依存関係の更新: go mod tidy
##
update-deps:
	@echo "==> Updating dependencies"
	@$(GOCMD) mod tidy

##
# Linux向けクロスコンパイル
##
build-linux:
	@echo "==> Cross compiling for Linux (amd64)"
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 $(MAKE) build
