package main

import (
	"context"
	"fmt"

	feed "github.com/bartick/learning-gRPC/go/proto-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	c := feed.NewFeedServiceClient(conn)

	in := new(emptypb.Empty)

	response, err := c.GetAllFeeds(context.Background(), in)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
