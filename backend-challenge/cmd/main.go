package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"

	challengeservice "github.com/Alexander021192/challenge/backend-challenge/internal/app"
	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
)

func main() {
	ctx := context.Background()

	// Start HTTP server // mux
	mux := runtime.NewServeMux()

	//newStore
	store, err := storage.NewStorage()
	if err != nil {
		log.Fatal(ctx, err)
	}

	// new TestServer
	testApiServer := challengeservice.NewTestApiServer(*store)

	// test
	fmt.Println(testApiServer.Echo(context.Background(), &pb.TestResponse{}))

	//register
	pb.RegisterTestApiHandlerServer(context.Background(), mux, testApiServer)
	// // http server

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	log.Fatalln(http.ListenAndServe("localhost:8080", handler))

	// Starting gRPC server
	// listner, err := net.Listen("tcp", "localhost:8080")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterTestApiServer(grpcServer, challenge.NewTestApiServer())

	// err = grpcServer.Serve(listner)
	// if err != nil {
	// 	log.Println(err)
	// }

}
