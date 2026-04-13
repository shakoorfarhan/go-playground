package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "whatsapp-chat/grpc-demo/chatpb"

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

	client := pb.NewChatServiceClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error starting chat: %v", err)
	}

	// receive messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				return
			}
			fmt.Printf("\n%s: %s\n", msg.GetUser(), msg.GetMessage())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Println("Start chatting! Type messages:")

	for scanner.Scan() {
		text := scanner.Text()

		if text == "exit" {
			break
		}

		err := stream.Send(&pb.ChatMessage{
			User:    username,
			Message: text,
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
	}
}
