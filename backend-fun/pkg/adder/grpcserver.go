package adder

type GRPCServer struct {
	func(s *GRPCServer) Add(ctx context.Context, req *api.Add)
}