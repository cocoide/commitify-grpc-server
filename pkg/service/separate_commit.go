package service

import (
	"context"
	"github.com/cocoide/commitify-grpc-server/pkg/entity"
	"github.com/cocoide/commitify-grpc-server/pkg/pb"
	"github.com/cocoide/commitify-grpc-server/pkg/usecase"
)

func NewSeparateCommitServer(su *usecase.SeparateCommitUsecase) pb.SeparateCommitServiceServer {
	return &separateCommitServer{su: su}
}

type separateCommitServer struct {
	pb.UnimplementedSeparateCommitServiceServer
	su *usecase.SeparateCommitUsecase
}

func (s *separateCommitServer) GenerateMultipleCommitMessage(ctx context.Context, req *pb.SeparateCommitRequest) (*pb.SeparateCommitResponse, error) {
	var fileChanges []entity.FileChange

	for _, v := range req.FileChanges {
		fileChanges = append(fileChanges, entity.ConvertPbToFileChange(v))
	}
	format := entity.ConvertPbCodeFormatToEntity(req.CodeFormat)
	language := entity.ConvertPbLanguageToEntity(req.Language)

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
