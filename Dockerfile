# syntax=docker/dockerfile:1.3-labs
FROM golang:1.17

RUN <<EOF
go install github.com/cosmtrek/air@latest
go install github.com/onsi/ginkgo/ginkgo@latest
go install github.com/rubenv/sql-migrate/sql-migrate@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
EOF
