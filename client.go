package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-demo/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRequest{Id: 1}

	res, err := client.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Name: %s\n", res.GetName())

	stream, err := client.WatchUsers(context.Background(), &pb.UserRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}

		fmt.Println("Stream:", msg.GetName())
	}
}
