.PHONY: api-up
api-up:
	docker-compose -f tools/docker-compose.yml up -d

goose-status:
	sh migrations/goose.sh status
	
goose-up:
	sh migrations/goose.sh up

goose-down:
	sh migrations/goose.sh down

## protoc
build-protoc:
	protoc --go_out=grpc \
	--go-grpc_out=grpc \
	grpc/proto/base_ball_api.proto

wire:
	wire ./...