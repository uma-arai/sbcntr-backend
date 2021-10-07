# sbcntr-backend 

The repository for sbcntr backend API application.

## Overview
This Repository provides demo web application coded by golang. This app is consist of three APIs.
- Healthcheck API
- Hello World API (turn inside the container)
- DB Access API(Notification, Item)

## Prerequisite
- Before start, you should build go runtime with go 1.16.
- Please clone this repository according to GOPATH env and move to this directory "sbcntr-backend".
  - You will find main directory in "$GOPATH/src/github.com/uma-arai/sbcntr-backend"
- Download required go module for this app by following command.
```bash
❯ go get golang.org/x/lint/golint
❯ go install 
❯ go mod download
```
- Set following env variables to connect DB.
  - DB_HOST
  - DB_USERNAME 
  - DB_PASSWORD 
  - DB_NAME

## How to Build and Deploy
### Golang
```bash
❯ make all
```
### Docker
```bash
❯ docker build -t sbcntr-backend:latest .
❯ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
sbcntr-backend                   latest              cdb20b70f267        58 minutes ago      4.45MB
:
❯ docker run -d -p 80:80 sbcntr-backend:latest
```

### REST API (after deploy)
```bash
❯ curl http://localhost:80/v1/helloworld
{"data":"Hello world"}

❯ curl http://localhost:80/healthcheck
null
```

## Options
- If you want to deploy this app with SSL signed by self, set env
variables TLS_CERT and TLS_KEY. 

## Notes
- We just check this operation only Mac OS (version 10.15).
