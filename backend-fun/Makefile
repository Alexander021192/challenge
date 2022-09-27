.PHONY: generate
generate:
		protoc --go_out=pkg --go_opt=paths=import \
		--go-grpc_out=pkg --go-grpc_opt=paths=import \
		proto/api/adder.proto;