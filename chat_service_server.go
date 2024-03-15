package main

import (
	"fmt"
	"io"
	"net"

	pb "github.com/anubhav100rao/url_shortner/proto"
	"google.golang.org/grpc"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

func NewChatServiceServer() *ChatServiceServer {
	return &ChatServiceServer{}
}

func (s *ChatServiceServer) Chat(stream pb.ChatService_ChatServer) error {
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println("Error receiving message")
			return err
		}

		fmt.Println("Message received: ", req.Message)
		fmt.Println("User: ", req.User)
		fmt.Println("\n\n")
		count++
		stream.Send(&pb.ChatMessage{Message: fmt.Sprintf("Message received %v", count), User: "Server"})
	}

	return nil
}

func SetUpChatServiceServer() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("GRPC server created")

	pb.RegisterChatServiceServer(grpcServer, NewChatServiceServer())
	fmt.Println("GRPC server registered")

	fmt.Println("Server started at :8081")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error serving")
		panic(err)
	}

}
