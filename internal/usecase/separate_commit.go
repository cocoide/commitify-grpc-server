package usecase

import (
	"github.com/cocoide/commitify-grpc-server/internal/entity"
	"github.com/cocoide/commitify-grpc-server/internal/gateway"
	"sort"
	"strings"
	"sync"
)

type SeparateCommitUsecase struct {
	og gateway.OpenAIGateway
	dg gateway.DeeplAPIGateway
	cu *CommitMessageUsecase
}

func NewSeparateCommitUsecaes(og gateway.OpenAIGateway, dg gateway.DeeplAPIGateway, cu *CommitMessageUsecase) *SeparateCommitUsecase {
	return &SeparateCommitUsecase{og: og, dg: dg, cu: cu}
}

func (u *SeparateCommitUsecase) GenerateMultipleFileMessages(changes []entity.FileChange, format entity.CodeFormatType, language entity.LanguageType) ([]entity.SeparatedCommitMessage, error) {
	var result []entity.SeparatedCommitMessage
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstError error

	for _, change := range changes {
		wg.Add(1)
		go func(change entity.FileChange) {
			defer wg.Done()

			var messages []string
			var err error
			code := u.generateCodePrompt(change)
			switch format {
			case entity.NormalFormat:
				messages, err = u.cu.GenerateNormalMessage(code, language)
			case entity.EmojiFormat:
				messages, err = u.cu.GenerateEmojiMessage(code, language)
			case entity.PrefixFormat:
				messages, err = u.cu.GeneratePrefixMessage(code, language)
			}
			if err != nil {
				mu.Lock()
				if firstError == nil {
					firstError = err
				}
				mu.Unlock()
				return
			}

			mu.Lock()
			result = append(result, entity.SeparatedCommitMessage{
				Messages:   messages,
				Filename:   change.Filename,
				ChangeType: change.ChangeType,
			})
			mu.Unlock()
		}(change)
	}

	if firstError != nil {
		return nil, firstError
	}
	wg.Wait()
	return result, nil
}

func (u *SeparateCommitUsecase) generateCodePrompt(change entity.FileChange) string {
	var builder strings.Builder

	builder.WriteString(change.Filename)
	builder.WriteString(", Diff: [")

	// Combine added and deleted lines into a single slice
	allLines := make([]entity.LineDiff, 0, len(change.CodeDiff.Added)+len(change.CodeDiff.Deleted))
	for _, lineDiff := range change.CodeDiff.Added {
		if lineDiff.Line != "" {
			allLines = append(allLines, entity.LineDiff{Index: lineDiff.Index, Line: "+ " + lineDiff.Line})
		}
	}
	for _, lineDiff := range change.CodeDiff.Deleted {
		if lineDiff.Line != "" {
			allLines = append(allLines, entity.LineDiff{Index: lineDiff.Index, Line: "- " + lineDiff.Line})
		}
	}

	// Sort combined lines by index
	sort.Slice(allLines, func(i, j int) bool {
		return allLines[i].Index < allLines[j].Index
	})

	// Write sorted lines to the builder
	for _, lineDiff := range allLines {
		builder.WriteString("\"")
		builder.WriteString(lineDiff.Line)
		builder.WriteString("\", ")
	}

	// Sort combined lines by index
	sort.Slice(allLines, func(i, j int) bool {
		return allLines[i].Index < allLines[j].Index
	})

	// Write sorted lines to the builder
	for i, lineDiff := range allLines {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString("\"")
		builder.WriteString(lineDiff.Line)
		builder.WriteString("\"")
	}

	builder.WriteString("]")

	return builder.String()
}
