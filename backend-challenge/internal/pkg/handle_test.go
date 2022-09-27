package challengeservice

import (
	"context"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.TestResponse) (*pb.TestResponse, error) {
	req.Id = 777
	req.Msg = "NEW" + req.Msg
	return req, nil
}
