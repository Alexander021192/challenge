package challengeservice

import (
	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

// NewTestApiServer return new instance TestApiServer
func NewTestApiServer() *testApiServer {
	return &testApiServer{}
}