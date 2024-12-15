run:
	go run ./cmd/main.go

gen.wire:
	wire ./cmd/factory

gen.proto:
	protoc \
	--proto_path=./pkg/service/proto \
	--go_opt=paths=source_relative \
	--go_out=./pkg/service/models \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=./pkg/service/models \
	./pkg/service/proto/*/*.proto