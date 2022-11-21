package challengeservice

import (
	"context"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

func GetPostById(ctx context.Context, store *storage.Storage, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	// get profileName, prodileImg bu sessionId
	_, err := store.FindBySession(req.SessionId)
	if err != nil {
		return &pb.GetPostResponse{Post: nil}, errSessionExpired
	}

	post, err := store.GetPostById(req.PostId)
	if err != nil {
		return &pb.GetPostResponse{Post: nil}, errStatusGetPosts
	}

	return &pb.GetPostResponse{Post: post}, nil
}
