api-up:
	docker-compose -f tools/docker-compose.yml up -d

api-delete:
	docker-compose -f tools/docker-compose.yml down --rmi all --volumes --remove-orphans

## protoc
build-protoc:
	protoc --go_out=grpc \
	--go-grpc_out=grpc \
	grpc/proto/base_ball_api.proto

wire:
	wire ./...