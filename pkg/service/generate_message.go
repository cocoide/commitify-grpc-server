package service

import (
	"context"

	"github.com/cocoide/commitify-grpc-server/pkg/pb"
)

func NewCommitMessage() pb.CommitMessageServiceServer {
	return &commitMessage{}
}

type commitMessage struct {
	pb.UnimplementedCommitMessageServiceServer
}

func (s commitMessage) GenerateCommitMessage(ctx context.Context, req *pb.CommitMessageRequest) (*pb.CommitMessageResponse, error) {
	return &pb.CommitMessageResponse{
		Messages: []string{"Feat A", "Add B"},
	}, nil
}
