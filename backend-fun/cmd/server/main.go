package main

import (
	"log"
	"net"

	"github.com/Alexander021192/just-for-fun/backend-fun/pkg/adder"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	s.RegisterService(&grpc.ServiceDesc{}, srv)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
