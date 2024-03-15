package main

import (
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/anubhav100rao/url_shortner/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		fmt.Println("We can establish connection at moment" +
			" either the server is down or the port is not open")
		panic(err)
	}

	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	fmt.Println("Client created")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.Chat(ctx)
	if err != nil {
		fmt.Println("Error in Chat in client side")
		panic(err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(
			&pb.ChatMessage{
				Message: fmt.Sprintf("Message %v", i),
				User:    "Client",
			},
		)

		if err != nil {
			log.Fatal(err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatal("CloseSend", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Recv", err)
		}
		fmt.Println("Message received: ", response.Message)
		fmt.Println("User: ", response.User)
		fmt.Println("\n\n")
	}
}
