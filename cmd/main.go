package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cocoide/commitify-grpc-server/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGenerateMessage() *generateMessage {
	return &generateMessage{}
}

type generateMessage struct {
	api.UnimplementedGenerateMessageServiceServer
}

func (s generateMessage) GenerateEnglishCommitMessage(ctx context.Context, req *api.GenerateMessageRequest) (*api.GenerateMessageResponse, error) {
	return &api.GenerateMessageResponse{
		Messages: []string{"Feat A", "Add B"},
	}, nil
}

func (s generateMessage) GenerateJapaneseCommitMessage(ctx context.Context, req *api.GenerateMessageRequest) (*api.GenerateMessageResponse, error) {
	return &api.GenerateMessageResponse{
		Messages: []string{"機能Aを開発", "機能Bを追加"},
	}, nil
}

func (s generateMessage) GeneratePrefixFormatCommitMessage(ctx context.Context, req *api.GenerateMessageRequest) (*api.GenerateMessageResponse, error) {
	return &api.GenerateMessageResponse{
		Messages: []string{"feat: A", "add: B"},
	}, nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Starting gRPC server on port: %v", port)
	api.RegisterGenerateMessageServiceServer(s, NewGenerateMessage())

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
