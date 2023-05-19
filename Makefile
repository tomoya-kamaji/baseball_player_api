help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'


api-up: # api-server起動
	swag init -g cmd/main.go 
	docker-compose -f docs/docker-compose.yml up -d
	docker-compose -f tools/docker-compose.yml up -d

.PHONY: goose-create
goose-create: # マイグレーションファイル作成 ex) make goose-create name=new_migration_name
	MIGRATION_FILE=$(name) sh migrations/goose.sh create

# dev,stg,pro
goose-up: # マイグレーション実行
	sh migrations/goose.sh up

goose-down: # マイグレーションロールバック
	sh migrations/goose.sh down

goose-status: # マイグレーションステータス
	sh migrations/goose.sh status

# test
goose-up-test: # マイグレーション実行
	sh migrations/goose_test.sh up

goose-down-test: # マイグレーションロールバック
	sh migrations/goose_test.sh down

goose-status-test: # マイグレーションステータス
	sh migrations/goose_test.sh status

build-protoc: # protoファイルからgoファイルを生成
	protoc --go_out=grpc \
	--go-grpc_out=grpc \
	grpc/proto/base_ball_api.proto

wire: # wireで依存関係を解決
	wire ./...

.PHONY: lint-all
lint-all: imports lint vet vet-for-loop

.PHONY: imports
imports:
	(! goimports -d $(shell find . -type f -name '*.go' -type f -not -name 'wire_gen.go' -not -name '*.pb.go' -type f -not -path "*/mock/*") | grep '^')

.PHONY: lint # lint
lint:
	staticcheck ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: vet-for-loop
vet-for-loop:
	go vet -vettool=`which itervar` ./...
