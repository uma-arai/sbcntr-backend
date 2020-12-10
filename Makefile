# Golang parameter(gofmt: code formatter, go vet/golist: lint tools)
GOCMD=go
GOFMT=go fmt
GOVET=$(GOCMD) vet
GOLINT=golint
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

BUILD_PATH=.
APP_PATH=.

all:
	make validate
	make build
	make run

# delete debug symbols by golang option -dlflags "-s -w"
build:
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -ldflags "-s -w" $(APP_PATH)/main.go
validate:
	$(GOFMT) ./...
	$(GOVET) ./...; $(GOLINT) -set_exit_status ./...
	$(GOTEST) -v ./...
clean:
	rm -f $(BUILD_PATH)/$(BINARY_NAME)
	rm -f $(BUILD_PATH)/$(BINARY_UNIX)
run:
	$(BUILD_PATH)/$(BINARY_NAME)

# cross-compile for linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

