# setup

```
docker-compose build
docker-compose run app sql-migrate up
docker-compose up
```

# test

```
docker-compose up -d mysql
docker-compose exec mysql mysql -e "create database go_api_tut_test"
docker-compose run app sql-migrate up -env test
docker-compose run app ginkgo -r
```

# lint

```
docker-compose run app golangci-lint run
```
