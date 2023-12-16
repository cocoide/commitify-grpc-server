package handler

import (
	"context"
	"fmt"
	"github.com/cocoide/commitify-grpc-server/internal/domain/entity"
	"github.com/cocoide/commitify-grpc-server/internal/usecase"

	pb "github.com/cocoide/commitify-grpc-server/pkg/grpc"
)

func NewCommitMessageServiceServer(cu *usecase.CommitMessageUsecase) pb.CommitMessageServiceServer {
	return &commitMessageServiceServer{cu: cu}
}

type commitMessageServiceServer struct {
	pb.UnimplementedCommitMessageServiceServer
	cu *usecase.CommitMessageUsecase
}

func (s commitMessageServiceServer) GenerateCommitMessage(ctx context.Context, req *pb.CommitMessageRequest) (*pb.CommitMessageResponse, error) {
	var messages []string
	var err error
	codeFormat := entity.ConvertPbCodeFormatToEntity(req.CodeFormat)
	language := entity.ConvertPbLanguageToEntity(req.Language)
	switch codeFormat {
	case entity.PrefixFormat:
		messages, err = s.cu.GeneratePrefixMessage(req.InputCode, language)
		if err != nil {
			return nil, err
		}
	case entity.EmojiFormat:
		messages, err = s.cu.GenerateEmojiMessage(req.InputCode, language)
		if err != nil {
			return nil, err
		}
	case entity.NormalFormat:
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
