package service

import (
	"context"
	"fmt"

	"github.com/cocoide/commitify-grpc-server/pkg/enum"
	"github.com/cocoide/commitify-grpc-server/pkg/pb"
)

func NewCommitMessage() pb.CommitMessageServiceServer {
	return &commitMessage{}
}

type commitMessage struct {
	pb.UnimplementedCommitMessageServiceServer
}

func (s commitMessage) GenerateCommitMessage(ctx context.Context, req *pb.CommitMessageRequest) (*pb.CommitMessageResponse, error) {
	var messages []string
	// messagesを生成する処理はUsecase層で行う予定
	switch req.Language {
	case enum.ENGLISH:
		switch req.CodeFormat {
		case enum.PREFIX_FORMAT:
			messages = []string{"feat: A", "fix: B"}
		case enum.EMOJI_FORMAT:
			messages = []string{"🎉 feat A", "🐛 fix B"}
		case enum.NORMAL_FORMAT:
			messages = []string{"Feat A", "Add B"}
		default:
			return nil, fmt.Errorf("CodeFormat not input")
		}
	case enum.JAPANESE:
		switch req.CodeFormat {
		case enum.EMOJI_FORMAT:
			messages = []string{"🎉  Aのリリース", "🐛 Bのバグ修正"}
		case enum.NORMAL_FORMAT:
			messages = []string{"Aのリリース", "Bのバグ修正"}
		case enum.PREFIX_FORMAT:
			return nil, fmt.Errorf("PREFIX_FORMAT not supported in JAPANESE")
		default:
			return nil, fmt.Errorf("CodeFormat not input")
		}
	default:
		return nil, fmt.Errorf("Language not input")
	}
	return &pb.CommitMessageResponse{
		Messages: messages,
	}, nil
}
