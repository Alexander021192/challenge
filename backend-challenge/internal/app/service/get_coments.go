package challengeservice

import (
	"context"
	"errors"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errStatusGetComment = errors.New("failed get comments")
)

func GetComments(ctx context.Context, store *storage.Storage, req *pb.EmptyRequest) (*pb.GetCommentsResponse, error) {
	// get profileName, prodileImg bu sessionId
	_, err := store.FindBySession(req.SessionId)
	if err != nil {
		return &pb.GetCommentsResponse{Comments: nil}, errSessionExpired
	}

	comments, err := store.GetComments()
	if err != nil {
		return &pb.GetCommentsResponse{Comments: nil}, errStatusGetComment
	}

	return &pb.GetCommentsResponse{Comments: comments}, nil
}
