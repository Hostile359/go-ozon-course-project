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

## Run
Run server

```make run```

Run test client

```make grpc_client```
