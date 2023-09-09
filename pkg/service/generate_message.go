package service

import (
	"context"

	"github.com/cocoide/commitify-grpc-server/pkg/pb"
)

func NewGenerateMessage() pb.GenerateMessageServiceServer {
	return &generateMessage{}
}

type generateMessage struct {
	pb.UnimplementedGenerateMessageServiceServer
}

func (s generateMessage) GenerateEnglishCommitMessage(ctx context.Context, req *pb.GenerateMessageRequest) (*pb.GenerateMessageResponse, error) {
	return &pb.GenerateMessageResponse{
		Messages: []string{"Feat A", "Add B"},
	}, nil
}

func (s generateMessage) GenerateJapaneseCommitMessage(ctx context.Context, req *pb.GenerateMessageRequest) (*pb.GenerateMessageResponse, error) {
	return &pb.GenerateMessageResponse{
		Messages: []string{"機能Aを開発", "機能Bを追加"},
	}, nil
}

func (s generateMessage) GeneratePrefixFormatCommitMessage(ctx context.Context, req *pb.GenerateMessageRequest) (*pb.GenerateMessageResponse, error) {
	return &pb.GenerateMessageResponse{
		Messages: []string{"feat: A", "add: B"},
	}, nil
}
