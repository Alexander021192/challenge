create: clean
		protoc --proto_path=proto proto/*.proto --go_out=gen/proto/
		protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/proto/
		protoc -I ./proto --grpc-gateway_out ./gen/proto/challenge/ \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			proto/challenge.proto

clean:
	rm gen/proto/challenge/*.go

# export PATH="$PATH:$(go env GOPATH)/bin"