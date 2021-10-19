# setup

```
docker-compose run app go install github.com/cosmtrek/air@latest
docker-compose run app go install github.com/rubenv/sql-migrate/sql-migrate@latest
docker-compose run app sql-migrate up
docker-compose up
```

# test

```
docker-compose up -d mysql
docker-compose exec mysql mysql -e "create database go_api_tut_test"
docker-compose run app sql-migrate up -env test
docker-compose run app go install github.com/onsi/ginkgo/ginkgo@latest
docker-compose run app ginkgo -r
```

# lint

```
docker-compose run app bash -c 'curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1'
docker-compose run app golangci-lint run
```
