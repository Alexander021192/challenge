package challengeservice

import (
	"context"
	"errors"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errStatusGetPosts = errors.New("failed get posts")
)

func GetPosts(ctx context.Context, store *storage.Storage, req *pb.EmptyRequest) (*pb.GetPostsResponse, error) {
	// get profileName, prodileImg bu sessionId
	_, err := store.FindBySession(req.SessionId)
	if err != nil {
		return &pb.GetPostsResponse{Posts: nil}, errSessionExpired
	}

	posts, err := store.GetPosts()
	if err != nil {
		return &pb.GetPostsResponse{Posts: nil}, errStatusGetPosts
	}

	return &pb.GetPostsResponse{Posts: posts}, nil
}
