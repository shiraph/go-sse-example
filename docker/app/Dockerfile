FROM golang:1.20.1-alpine3.17

ENV GOOS linux
ENV GOARCH amd64
ENV PATH $PATH:/go/bin/linux_amd64
ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app
