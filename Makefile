CMD_DIR=./cmd/

.PHONY: proto-all
proto-all:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
    api/protobuf/*.proto

.PHONY: proto-health
proto-health:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	api/protobuf/health.proto

run:
	@set -a; . ./.env; set +a; go run $(CMD_DIR)main.go

