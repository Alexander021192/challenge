package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	challenge "github.com/Alexander021192/challenge/backend-challenge/internal/pkg"
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
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// new TestServer
		testApiServer := challenge.NewTestApiServer()

		//register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, testApiServer)
		// // http server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, challenge.NewTestApiServer())

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}

}
