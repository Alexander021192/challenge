package adder

import (
	"context"

	"github.com/Alexander021192/just-for-fun/backend-fun/pkg/api"
)

// GRPCServer...
type GRPCServer struct{}

func(s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	return &api.AddResponce{Res: req.GetX() + req.GetY()}, nil
}