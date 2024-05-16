gen-pb:
	@protoc \
		--proto_path=protobuf 							"protobuf/profit.proto" \
		--go_out=internal/api/grpc/service/profit 		--go_opt=paths=source_relative \
  		--go-grpc_out=internal/api/grpc/service/profit 	--go-grpc_opt=paths=source_relative

test:
	@go test -coverprofile=coverage.out ./...

run-rpc:
	@go run ./cmd/main/main.go rpc

run-http:
	@go run ./cmd/main/main.go http

# generate jwt token with arg of hours
# if negative, expired token will get generated
gen-tkn:
	@go run ./scripts/main/generate_jwt_token.go $(hours)
