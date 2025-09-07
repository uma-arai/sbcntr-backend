# sbcntr-backend

Backend API repository for the book "AWS Container Design and Construction [Practical] Introduction 2nd Edition".

## Overview

This is a Golang-based API server using the echo framework.
Among the many frameworks available for Golang, echo was chosen for its comprehensive features for implementing REST API servers and its excellent documentation.

The connection between the API server and DB (Postgres) uses sqlx[^sqlx], an O/R mapper library.

[^sqlx]: <https://jmoiron.github.io/sqlx/>

The backend application provides two services with `/v1` prefix for all API endpoints:

1. Pet Service (`/pets`)
   - `GET /pets` - Get pet list
   - `POST /pets/:id/like` - Like/unlike a pet
   - `POST /pets/:id/reservation` - Reserve a pet

2. Notification Service (`/notifications`)
   - `GET /notifications` - Get notification list
   - `POST /notifications/read` - Mark notifications as read

## Intended Use

Please use this repository according to the book's content.

## Local Usage

### Prerequisites

- Go version 1.23.x is required.
- Clone this repository to an appropriate directory according to your GOPATH location.
- Download modules using the following commands:

```bash
go get golang.org/x/lint/golint
go install
go mod download
```

- This backend API requires DB connection. Set the following environment variables:
  - DB_HOST
  - DB_USERNAME
  - DB_PASSWORD
  - DB_NAME

### Database Setup

Start a local Postgres server beforehand.

### Build & Deploy

#### Running Locally

```text
export DB_HOST=localhost
export DB_USERNAME=sbcntrapp
export DB_PASSWORD=password
export DB_NAME=sbcntrapp
export DB_CONN=1
```

```bash
make all
```

#### Running with Docker

```bash
$ docker build -t sbcntr-backend:latest .
$ docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
sbcntr-backend                   latest              cdb20b70f267        58 minutes ago      4.45MB
:
$ docker run -d -p 80:80 sbcntr-backend:latest
```

### Connectivity Check After Deployment

```bash
$ curl http://localhost:80/v1/helloworld
{"data":"Hello world"}

$ curl http://localhost:80/healthcheck
null
```

## Notes

- Operation has been verified only on Mac OS Sequoia 15.6.
