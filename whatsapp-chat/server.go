package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"

	pb "whatsapp-chat/grpc-demo/chatpb"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

type server struct {
	pb.UnimplementedChatServiceServer

	mu      sync.Mutex
	clients map[string]pb.ChatService_ChatServer
}

func newServer() *server {
	return &server{
		clients: make(map[string]pb.ChatService_ChatServer),
	}
}

func startSubscriber(s *server) {
	pubsub := rdb.Subscribe(ctx, "chat-room-1")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			log.Println("Redis broadcast:", msg.Payload)

			s.mu.Lock()
			for _, stream := range s.clients {
				stream.Send(&pb.ChatMessage{
					User:    "broadcast",
					Message: msg.Payload,
				})
			}
			s.mu.Unlock()
		}
	}()
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {

	var user string

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if user == "" {
			user = msg.User

			s.mu.Lock()
			s.clients[user] = stream
			s.mu.Unlock()
		}

		log.Printf("%s: %s", msg.User, msg.Message)

		err = rdb.Publish(ctx, "chat-room-1", msg.Message).Err()
		if err != nil {
			log.Printf("Redis error: %v", err)
		}

		// local broadcast
		s.mu.Lock()
		for u, clientStream := range s.clients {
			if u == user {
				continue
			}

			clientStream.Send(&pb.ChatMessage{
				User:    msg.User,
				Message: msg.Message,
			})
		}
		s.mu.Unlock()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := newServer()

	startSubscriber(s) // 🔥 IMPORTANT

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, s)

	log.Println("Chat server running on :50051")
	grpcServer.Serve(lis)
}
