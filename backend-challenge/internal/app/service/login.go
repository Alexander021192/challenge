package challengeservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
	"golang.org/x/crypto/bcrypt"
)

var (
	errIncorectEmailOrPassword   = errors.New("incorrect email or password")
	errNotAuthenticated          = errors.New("not authenticated")
	errStatusInternalServerError = errors.New("Internal Server Error")
)

func Login(ctx context.Context, store *storage.Storage, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	// find user and comapare password
	u, err := store.FindByEmail(req.Email)
	fmt.Println(u, err)
	if err != nil || !ComparePassword(req.Password, u.EncryptedPassword) {
		return &pb.LoginResponse{SessionId: ""}, errIncorectEmailOrPassword
	}

	//generate sessionId and save it for 2 hours (gorouting with deadline)
	sessionId, err := store.SessionSave(u.Email)
	fmt.Println(sessionId, err)
	if err != nil {
		return &pb.LoginResponse{SessionId: ""}, errStatusInternalServerError
	}

	return &pb.LoginResponse{SessionId: sessionId}, nil
}

func ComparePassword(password string, userPassword string) bool {
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)) == nil)
	return userPassword == password
}
