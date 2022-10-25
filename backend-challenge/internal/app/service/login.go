package challengeservice

import (
	"context"
	"errors"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errIncorectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated        = errors.New("not authenticated")
	errStatusInternalServerError = errors.New("Internal Server Error")
)

func Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var sessionId string

	// find user and comapare password
	u, err := s.store.User().FindByEmail(req.Email)
	if err != nil || !u.ComparePassword(req.Password) {
		return &pb.LoginResponse{SessionId: ""}, errIncorectEmailOrPassword
	}

	//generate sessionId and save it for 2 hours (gorouting with deadline)
	if sessionId, err := s.SessionSave(u.Id); err != nil {
		return &pb.LoginResponse{SessionId: ""}, errStatusInternalServerError
	}

	return &pb.LoginResponse{SessionId: sessionId}, nil
}
