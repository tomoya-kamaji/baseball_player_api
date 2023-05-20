# Onion Architecture for Golang

## アプリケーション立ち上げ
```
// データベースのマイグレーション
make goose-up

// アプリケーションの立ち上げ
make api-up
```

## SwaggerでAPI仕様書を作成＆閲覧
```
make gen-swagger
```
閲覧する場合は以下のURLにアクセス(Swagger UI)
http://localhost:8002


## アプリケーション概要
- 野球選手の成績を管理、閲覧する


## ディレクトリ構成

```
├── Makefile
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── grpc
│   ├── base_ball_api.pb.go
│   ├── base_ball_api_grpc.pb.go
│   └── proto
│       └── base_ball_api.proto
├── migrations
│   └── goose.sh
├── src
│   ├── adapter
│   │   └── gorm
│   │       ├── config.go
│   │       └── gorm.go
│   ├── api
│   │   └── api.go
│   ├── domain
│   │   └── player
│   │       ├── player.go
│   │       └── playerRepository.go
│   ├── infrastructure
│   │   ├── entity
│   │   │   ├── baseEntity.go
│   │   │   └── playerEntity.go
│   │   ├── repositoryImpl
│   │   │   └── playerRepositoryImpl.go
│   │   └── transactionImpl
│   │       └── transactionManagerImpl.go
│   ├── usecase
│   │   ├── player
│   │   │   ├── createPlayerUsecace.go
│   │   │   └── fetchPlayerUsecace.go
│   │   └── transaction
│   │       └── transactionManeger.go
│   └── util
│       └── sliceutil
│           ├── slice.go
│           └── slice_test.go
└── tools
    ├── api
    │   ├── Dockerfile
    │   └── modd.conf
    ├── data
    │   └── redis
    │       └── dump.rdb
    └── docker-compose.yml

```
## セットアップ

Dockerを立ち上げる
```
make api-up
```

