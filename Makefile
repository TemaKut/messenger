install.bin:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ./bin v1.63.4

run.apigateway: 
	go run ./internal/services/apigateway/cmd

run.auth: 
	go run ./internal/services/auth/cmd

gen.wire:
	wire ./internal/services/apigateway/cmd/factory
	wire ./internal/services/auth/cmd/factory

gen.svc_proto:
	protoc \
	--proto_path=./pkg/proto/service/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/proto/service/gen \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/proto/service/gen \
	./pkg/proto/service/proto/*/*.proto

gen.client_proto:
	protoc \
	--proto_path=./pkg/proto/client/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/proto/client/gen \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/proto/client/gen \
	./pkg/proto/client/proto/*/*.proto

gen.proto:
	@make gen.svc_proto
	@make gen.client_proto

lint:
	./bin/golangci-lint run ./...
