generate:
	cd api; buf generate

run:
	go run cmd/main.go

migrate:
	migrate -path migrations -database "postgres://localhost/challenge_test?sslmode=disable" up

clean:
	rm -rf pkg

# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# for export path env
# export PATH="$PATH:$(go env GOPATH)/bin"


# for show database 
# psql -d challenge_test
