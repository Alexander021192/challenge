package adder

import (
	""
	"context"
)

// GRPCServer...
type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponce, error) {
	return &api.AddResponce{Result: req.GetX() + req.GetY()}, nil
}
