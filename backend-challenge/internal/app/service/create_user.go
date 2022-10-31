package challengeservice

import (
	"context"
	"errors"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errStatusUnprocessableEntity = errors.New("failed create user")
)

func CreateUser(ctx context.Context, store *storage.Storage, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	u := &storage.User{
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := store.Create(u)
	if err != nil {
		return &pb.CreateUserResponse{UserId: 0}, errStatusUnprocessableEntity
	}

	// delete true password
	u.Password = ""
	return &pb.CreateUserResponse{UserId: id}, nil
}
