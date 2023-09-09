package service

import (
	"context"
	"fmt"

	"github.com/cocoide/commitify-grpc-server/pkg/enum"
	"github.com/cocoide/commitify-grpc-server/pkg/gateway"
	"github.com/cocoide/commitify-grpc-server/pkg/pb"
	"github.com/cocoide/commitify-grpc-server/pkg/usecase"
)

func NewCommitMessageServiceServer(og gateway.OpenAIGateway, dg gateway.DeeplAPIGateway, cu usecase.CommitMessageUseCase) pb.CommitMessageServiceServer {
	return &commitMessageServiceServer{og: og, dg: dg, cu: cu}
}

type commitMessageServiceServer struct {
	pb.UnimplementedCommitMessageServiceServer
	og gateway.OpenAIGateway
	dg gateway.DeeplAPIGateway
	cu usecase.CommitMessageUseCase
}

func (s commitMessageServiceServer) GenerateCommitMessage(ctx context.Context, req *pb.CommitMessageRequest) (*pb.CommitMessageResponse, error) {
	var messages []string
	var err error
	codeFormat := enum.ConvertPbCodeFormat(req.CodeFormat)
	language := enum.ConvertPbLanguage(req.Language)
	switch codeFormat {
	case enum.PrefixFormat:
		messages, err = s.cu.GeneratePrefixMessage(req.InputCode, language)
		if err != nil {
			return nil, err
		}
	case enum.EmojiFormat:
		messages, err = s.cu.GenerateEmojiMessage(req.InputCode, language)
		if err != nil {
			return nil, err
		}
	case enum.NormalFormat:
		messages, err = s.cu.GenerateNormalMessage(req.InputCode, language)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("CodeFormat not input")
	}
	return &pb.CommitMessageResponse{
		Messages: messages,
	}, nil
}
