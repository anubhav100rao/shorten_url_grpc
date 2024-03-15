package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/anubhav100rao/url_shortner/proto"
	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUrlShortnerServer
	mapUrls map[string]string
}

func NewServer() *server {
	mapUrls := make(map[string]string)
	return &server{mapUrls: mapUrls}
}

func (s *server) ShortenUrl(ctx context.Context, req *pb.ShortenUrlRequest) (*pb.ShortenUrlResponse, error) {
	uri := req.Url
	if uri == "" {
		return nil, nil
	}

	// reduce the length of the uri
	uri = uri[:5]
	// store the uri in the map
	s.mapUrls[uri] = uri

	return &pb.ShortenUrlResponse{ShortUrl: uri}, nil
}

func (s *server) GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	shortUrl := req.ShortUrl
	if shortUrl == "" {
		return nil, nil
	}
	// get the uri from the map
	uri := s.mapUrls[shortUrl]
	return &pb.GetUrlResponse{Url: uri}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("GRPC server created")

	pb.RegisterUrlShortnerServer(grpcServer, NewServer())
	fmt.Println("GRPC server registered")

	fmt.Println("Server started at :8080")

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error starting server")
		panic(err)
	}

}
