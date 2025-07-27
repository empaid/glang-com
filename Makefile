build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

run-orders:
	@go run services/orders/*.go
gen:
	@protoc \
		--proto_path=protobuf \
		--go_out=paths=source_relative:services/common/genproto/orders \
		--go-grpc_out=paths=source_relative:services/common/genproto/orders \
		protobuf/orders.proto