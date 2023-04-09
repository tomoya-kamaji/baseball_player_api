.PHONY: api-up
api-up:
	docker-compose -f tools/docker-compose.yml up -d

## protoc
build-protoc:
	protoc --go_out=grpc \
	--go-grpc_out=grpc \
	grpc/proto/base_ball_api.proto