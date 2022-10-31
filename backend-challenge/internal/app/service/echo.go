package challengeservice

import (
	"context"

	pb "github.com/Alexander021192/challenge/backend-challenge/pkg"
)

func Echo(ctx context.Context, req *pb.TestResponse) (*pb.TestResponse, error) {
	req.Id = 1 + req.Id
	req.Msg = "HELLO WORLD" + req.Msg
	return req, nil
}
