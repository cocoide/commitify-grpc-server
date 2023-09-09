package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cocoide/commitify-grpc-server/pkg/gateway"
	"github.com/cocoide/commitify-grpc-server/pkg/pb"
	"github.com/cocoide/commitify-grpc-server/pkg/service"
	"github.com/cocoide/commitify-grpc-server/pkg/usecase"
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
	cu := usecase.NewCommitMessageUseCaes(og, dg)
	pb.RegisterCommitMessageServiceServer(s, service.NewCommitMessageServiceServer(og, dg, *cu))

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
