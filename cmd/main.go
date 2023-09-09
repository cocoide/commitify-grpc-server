package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cocoide/commitify-grpc-server/pkg/pb"
	"github.com/cocoide/commitify-grpc-server/pkg/service"
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
	log.Printf("Starting gRPC server on port: %v", port)
	pb.RegisterGenerateMessageServiceServer(s, service.NewGenerateMessage())

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
