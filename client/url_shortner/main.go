package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/anubhav100rao/url_shortner/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		fmt.Println(
			"We can establish connection at moment" +
				" either the server is down or the port is not open",
		)
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUrlShortnerClient(conn)
	fmt.Println("Client created")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.ShortenUrl(ctx, &pb.ShortenUrlRequest{Url: "https://www.google.com"})
	if err != nil {
		fmt.Println("Error in ShortenUrl")
		panic(err)
	}

	fmt.Println("Shortened URL: ", response.ShortUrl)

}
