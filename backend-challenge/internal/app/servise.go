package challengeservice

import (
	"context"

	service "github.com/Alexander021192/challenge/backend-challenge/internal/app/service"
	"github.com/Alexander021192/challenge/backend-challenge/internal/app/storage"
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
	storage *storage.Storage
}

// NewTestApiServer return new instance TestApiServer
func NewTestApiServer(storage storage.Storage) *testApiServer {
	return &testApiServer{
		storage: &storage,
	}
}

// Implemetations

// Echo implementation
func (s *testApiServer) Echo(ctx context.Context, req *pb.TestResponse) (*pb.TestResponse, error) {
	response, err := service.Echo(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Login implementation
func (s *testApiServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	response, err := service.Login(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Create  user implementation
func (s *testApiServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	response, err := service.CreateUser(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Create Comment implementation
func (s *testApiServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	response, err := service.CreateComment(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Create Comment implementation
func (s *testApiServer) GetComments(ctx context.Context, req *pb.EmptyRequest) (*pb.GetCommentsResponse, error) {
	response, err := service.GetComments(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Create Post implementation
func (s *testApiServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	response, err := service.CreatePost(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Create Comment implementation
func (s *testApiServer) GetPosts(ctx context.Context, req *pb.EmptyRequest) (*pb.GetPostsResponse, error) {
	response, err := service.GetPosts(ctx, s.storage, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}