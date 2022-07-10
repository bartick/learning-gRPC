package main

import (
	"fmt"
	"net"

	feed "github.com/bartick/learning-gRPC/go/proto-go"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println(err)
		return
	}

	feedServer := Server{}

	grpcServer := grpc.NewServer()

	feed.RegisterFeedServiceServer(grpcServer, &feedServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server started")

}
