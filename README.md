# setup

```
docker-compose build
docker-compose run app sql-migrate up
docker-compose up
```

# grpc

```
docker-compose up -d grpc
docker-compose run app go run pkg/grpc/client/main.go
```

# regenerate grpc code

```
docker-compose run app protoc --go_out . --go-grpc_out . pkg/grpc/proto/hello.proto
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
