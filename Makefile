gen-pb:
	@protoc \
		--proto_path=protobuf "protobuf/max_profit.proto" \
		--go_out=internal/proto --go_opt=paths=source_relative \
  	--go-grpc_out=internal/proto --go-grpc_opt=paths=source_relative
