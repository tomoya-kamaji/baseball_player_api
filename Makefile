.PHONY: api-up
api-up:
	docker-compose -f tools/docker-compose.yml up -d

## protoc
build-protoc:
	protoc --go-grpc_out=grpc --go-grpc_opt=require_unimplemented_servers=false proto/base_ball_server.proto