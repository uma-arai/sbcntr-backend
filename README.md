# cnappdemo (Demo Web App)

## Overview
This Repository provides demo web application coded by golang. This app is consist of three APIs.
- Healthcheck API
- Hello World API (turn inside the container)
- DB Access API

## Prerequisite
- Before start, you should build go runtime with go 1.13.5.
- Please clone this repository according to GOPATH env and move to this directory "cnappdemo".
  - You will find main directory in "$GOPATH/src/github.com/iselegant/cnappdemo"
- Download required go module for this app by following command.
```bash
❯ go get golang.org/x/lint/golint
❯ go get 
❯ go mod download
```
- Set following env variables to connect DB.
  - CNAPP_DB_HOST
  - CNAPP_DB_USERNAME 
  - CNAPP_DB_PASSWORD 
  - CNAPP_DB_NAME

## How to Build and Deploy
### Golang
```bash
❯ make all
```
### Docker
```bash
❯ docker build -t cnappdemo:latest .
❯ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
cnappdemo                   latest              cdb20b70f267        58 minutes ago      4.45MB
:
❯ docker run -d -p 80:80 cnappdemo:latest
```

### REST API (after deploy)
```bash
❯ curl http://localhost:80/cnappdemo/v1/helloworld
"Hello world!"

❯ curl http://localhost:80/healthcheck
null
```

## Options
- If you want to deploy this app with SSL signed by self, set env
variables TLS_CERT and TLS_KEY. 

## Notes
- We just check this operation only Mac OS (version 10.15).
