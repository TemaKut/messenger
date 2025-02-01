run.apigateway: 
	go run ./internal/services/apigateway/cmd

run.auth: 
	go run ./internal/services/auth/cmd

gen.wire:
	wire ./internal/services/apigateway/cmd/factory
	wire ./internal/services/auth/cmd/factory

gen.proto:
	protoc \
	--proto_path=./pkg/proto/service/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/proto/service/gen \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/proto/service/gen \
	./pkg/proto/service/proto/*/*.proto