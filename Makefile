run.auth:
	go run ./internal/services/auth/cmd

gen.wire:
	wire ./internal/services/auth/cmd/factory

gen.service_proto:
	protoc \
	--proto_path=./pkg/service/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/service/models \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/service/models \
	./pkg/service/proto/*/*.proto