# setup

```
docker-compose run app go mod tidy
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
docker-compose run go install github.com/onsi/ginkgo/ginkgo@latest
docker-compose run ginkgo -r
```
