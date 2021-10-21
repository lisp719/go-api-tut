# syntax=docker/dockerfile:1.3-labs
FROM golang:1.17

RUN <<eot
apt update
apt install -y protobuf-compiler
rm -rf /var/lib/apt/lists/*
go install github.com/cosmtrek/air@latest
go install github.com/onsi/ginkgo/ginkgo@latest
go install github.com/rubenv/sql-migrate/sql-migrate@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
eot
