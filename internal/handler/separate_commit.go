package handler

import (
	"context"
	entity2 "github.com/cocoide/commitify-grpc-server/internal/domain/entity"
	"github.com/cocoide/commitify-grpc-server/internal/usecase"
	pb "github.com/cocoide/commitify-grpc-server/pkg/grpc"
)

func NewSeparateCommitServer(su *usecase.SeparateCommitUsecase) pb.SeparateCommitServiceServer {
	return &separateCommitServer{su: su}
}

type separateCommitServer struct {
	pb.UnimplementedSeparateCommitServiceServer
	su *usecase.SeparateCommitUsecase
}

func (s *separateCommitServer) GenerateMultipleCommitMessage(ctx context.Context, req *pb.SeparateCommitRequest) (*pb.SeparateCommitResponse, error) {
	var fileChanges []entity2.FileChange

	for _, v := range req.FileChanges {
		fileChanges = append(fileChanges, entity2.ConvertPbToFileChange(v))
	}
	format := entity2.ConvertPbCodeFormatToEntity(req.CodeFormat)
	language := entity2.ConvertPbLanguageToEntity(req.Language)

	commitMessages, err := s.su.GenerateMultipleFileMessages(fileChanges, format, language)
	if err != nil {
		return nil, err
	}
	var response pb.SeparateCommitResponse

	for _, v := range commitMessages {
		response.SeparatedCommits = append(response.SeparatedCommits, v.ConvertToPbCommitMessages())
	}
	return &response, nil
}
