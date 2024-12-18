run:
	go run ./internal/services/auth/cmd

gen.wire:
	wire ./internal/services/auth/cmd/factory
	wire ./internal/services/apigateway/cmd/factory

gen.proto:
	protoc \
	--proto_path=./pkg/service/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/service/models/proto \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/service/models/proto \
	./pkg/service/proto/*/*.proto