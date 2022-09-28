package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	challengeservice "github.com/Alexander021192/challenge/backend-challenge/internal/app"
)

// type testApiServer struct {
// 	pb.UnimplementedTestApiServer
// }

// func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
// 	return &pb.UserResponse{}, nil
// }

// func (s *testApiServer) Echo(ctx context.Context, req *pb.TestResponse) (*pb.TestResponse, error) {
// 	req.Id = 777
// 	req.Msg = "NEW" + req.Msg
// 	return req, nil
// }

func main() {

	// Start HTTP server
	// mux
	mux := runtime.NewServeMux()

	// new TestServer
	testApiServer := challengeservice.NewTestApiServer()

	// test 
	fmt.Println(testApiServer.Echo(context.Background(), &pb.TestResponse{}))
	fmt.Println(testApiServer.GetUser(context.Background(), &pb.UserRequest{Uuid: "testGetUser"}))

	//register
	pb.RegisterTestApiHandlerServer(context.Background(), mux, testApiServer)
	// // http server
	log.Fatalln(http.ListenAndServe("localhost:8080", mux))


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
