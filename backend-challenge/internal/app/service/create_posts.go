package challengeservice

import (
	"context"
	"errors"

	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

var (
	errStatusCreatePost = errors.New("failed create comment")
)

func CreatePost(ctx context.Context, store *storage.Storage, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	// get profileName, prodileImg bu sessionId
	u, err := store.FindBySession(req.SessionId)
	if err != nil {
		return &pb.CreatePostResponse{PostId: 0}, errSessionExpired
	}
	
	// need replace this interface for protobuf interface
	//challenge.Posts *challenge.Comment
	post := &storage.Post{
		Author: u.ProfileName,
		Title: req.Title,
		Location: req.Location,
		PostText: req.PostText,
		PostImg: req.PostImg,
	}

	id, err := store.CreatePost(post)
	if err != nil {
		return &pb.CreatePostResponse{PostId: 0}, errStatusCreatePost
	}

	return  &pb.CreatePostResponse{PostId: id}, nil
}
