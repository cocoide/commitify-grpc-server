package main

import (
	"context"
	"fmt"
	"github.com/cocoide/commitify-grpc-server/internal/gateway"
	"github.com/cocoide/commitify-grpc-server/internal/handler"
	"github.com/cocoide/commitify-grpc-server/internal/usecase"
	"log"
	"net"

	pb "github.com/cocoide/commitify-grpc-server/pkg/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}
	log.Printf("Starting gRPC server on port: %v", port)
	ctx := context.Background()
	og := gateway.NewOpenAIGateway(ctx)
	dg := gateway.NewDeeplAPIGateway()
	cu := usecase.NewCommitMessageUsecaes(og, dg)
	su := usecase.NewSeparateCommitUsecaes(og, dg, cu)
	pb.RegisterCommitMessageServiceServer(s, handler.NewCommitMessageServiceServer(cu))
	pb.RegisterSeparateCommitServiceServer(s, handler.NewSeparateCommitServer(su))

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
