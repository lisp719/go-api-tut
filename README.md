# setup

```
docker-compose run app go mod tidy
docker-compose run app go install github.com/cosmtrek/air@latest
docker-compose run app go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
docker-compose run app migrate -database "mysql://root:@tcp(mysql:3306)/go_api_tut_development?charset=utf8&parseTime=True&loc=Local" -path db/migrations/ up
docker-compose up
```

# test

```
docker-compose up -d mysql
docker-compose exec mysql mysql -e "create database go_api_tut_test"
docker-compose run app migrate -database "mysql://root:@tcp(mysql:3306)/go_api_tut_test?charset=utf8&parseTime=True&loc=Local" -path db/migrations/ up
docker-compose run go install github.com/onsi/ginkgo/ginkgo@latest
docker-compose run ginkgo -r
```
