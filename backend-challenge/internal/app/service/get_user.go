package challengeservice

import (
	"context"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

func GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{Name: "natali", Age: 38, Email: "example@test.com"}, nil
}
