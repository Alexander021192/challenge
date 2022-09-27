package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Alexander021192/challenge/backend-challenge/gen/proto/challenge"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.TestResponse{Msg: "Hello!BRO"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
