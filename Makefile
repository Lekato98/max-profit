gen-pb:
	@protoc \
		--proto_path=protobuf 					"protobuf/profit.proto" \
		--go_out=internal/api/grpc/profit 		--go_opt=paths=source_relative \
  		--go-grpc_out=internal/api/grpc/profit 	--go-grpc_opt=paths=source_relative
