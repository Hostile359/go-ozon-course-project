# Homework-1

## Prepare project

Install protoc

```make .deps```

Install buf

```brew install bufbuild/buf/buf```

Generate grpc api

```buf generate api```

To prepare project enviroment you need to set up your tg apikey:

```make prepare APIKEY="<your_tg_apikey>"```

## Prepare DB
Build and start docker compose

```make up-db```

Make migrations

```./migrate.sh```

Stop docker compose

```make down-db```

## Run
Run db service

```make run-db```

Run valid service

```make run-valid```

Run test client (out of date)

```make grpc_client```


## Tests
Run unit tests

```make .test```

Unit tests coverege

```make cover```

Run integration tests (you need to run db and valid services, before)

```make .integration_test```