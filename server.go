package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-demo/userpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received request for user ID: %d", req.GetId())
	// Simulate fetching user data based on the ID
	user := &pb.UserResponse{
		Name: fmt.Sprintf("User%d", req.GetId()),
	}
	return user, nil
}

func (s *server) WatchUsers(
	req *pb.UserRequest,
	stream pb.UserService_WatchUsersServer,
) error {

	log.Printf("Received request to watch users starting from ID: %d", req.GetId())

	for i := 0; i < 5; i++ {
		user := &pb.UserResponse{
			Name: fmt.Sprintf("User stream %d (base ID %d)", i, req.GetId()),
		}

		if err := stream.Send(user); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
