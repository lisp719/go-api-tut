# setup

```
docker-compose build
docker-compose run app sql-migrate up
docker-compose run grpc npm i
docker-compose up
```

# regenerate grpc code

```
docker-compose run app protoc --go_out . --go-grpc_out . pkg/proto/hello.proto
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
