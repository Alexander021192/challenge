package challengeservice

import (
	"context"

	service "github.com/Alexander021192/challenge/backend-challenge/internal/app/service"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

// NewTestApiServer return new instance TestApiServer
func NewTestApiServer() *testApiServer {
	return &testApiServer{}
}

// Implemetations

// Get user implementation
func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	response, err := service.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Echo implementation
func (s *testApiServer) Echo(ctx context.Context, req *pb.TestResponse) (*pb.TestResponse, error) {
	response, err := service.Echo(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Login implementation
func (s *testApiServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	response, err := service.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
