package challengeservice

import (
	"context"
	"errors"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errStatusCreateComment = errors.New("failed create comment")
	errSessionExpired = errors.New("failed create comment: SessionExpired")
)

func CreateComment(ctx context.Context, store *storage.Storage, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	// get profileName, prodileImg bu sessionId
	u, err := store.FindBySession(req.SessionId)
	if err != nil {
		return &pb.CreateCommentResponse{CommentId: 0}, errSessionExpired
	}
	
	comment := &storage.Comment{
		Author: u.ProfileName,
		ProfileImg: u.ProfileImg,
		Date: req.Date,
		Comment: req.Comment,
		CommentImg: req.CommentImg,
	}

	id, err := store.CreateComment(comment)
	if err != nil {
		return &pb.CreateCommentResponse{CommentId: 0}, errStatusCreateComment
	}

	return  &pb.CreateCommentResponse{CommentId: id}, nil
}
