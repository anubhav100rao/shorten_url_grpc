package main

import pb "github.com/anubhav100rao/url_shortner/proto"

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

func NewChatServiceServer() *ChatServiceServer {
	return &ChatServiceServer{}
}

func setUpChatServiceServer() {
}
