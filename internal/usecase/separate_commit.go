package usecase

import (
	entity2 "github.com/cocoide/commitify-grpc-server/internal/domain/entity"
	"github.com/cocoide/commitify-grpc-server/internal/domain/service"
	"sort"
	"strings"
	"sync"
)

type SeparateCommitUsecase struct {
	nlp  service.NLPService
	lang service.LangService
	cu   *CommitMessageUsecase
}

func NewSeparateCommitUsecaes(nlp service.NLPService, lang service.LangService, cu *CommitMessageUsecase) *SeparateCommitUsecase {
	return &SeparateCommitUsecase{nlp: nlp, lang: lang, cu: cu}
}

func (u *SeparateCommitUsecase) GenerateMultipleFileMessages(changes []entity2.FileChange, format entity2.CodeFormatType, language entity2.LanguageType) ([]entity2.SeparatedCommitMessage, error) {
	var result []entity2.SeparatedCommitMessage
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstError error

	for _, change := range changes {
		wg.Add(1)
		go func(change entity2.FileChange) {
			defer wg.Done()

			var messages []string
			var err error
			code := u.generateCodePrompt(change)
			switch format {
			case entity2.NormalFormat:
				messages, err = u.cu.GenerateNormalMessage(code, language)
			case entity2.EmojiFormat:
				messages, err = u.cu.GenerateEmojiMessage(code, language)
			case entity2.PrefixFormat:
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
			result = append(result, entity2.SeparatedCommitMessage{
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

func (u *SeparateCommitUsecase) generateCodePrompt(change entity2.FileChange) string {
	var builder strings.Builder

	builder.WriteString(change.Filename)
	builder.WriteString(", Diff: [")

	// Combine added and deleted lines into a single slice
	allLines := make([]entity2.LineDiff, 0, len(change.CodeDiff.Added)+len(change.CodeDiff.Deleted))
	for _, lineDiff := range change.CodeDiff.Added {
		if lineDiff.Line != "" {
			allLines = append(allLines, entity2.LineDiff{Index: lineDiff.Index, Line: "+ " + lineDiff.Line})
		}
	}
	for _, lineDiff := range change.CodeDiff.Deleted {
		if lineDiff.Line != "" {
			allLines = append(allLines, entity2.LineDiff{Index: lineDiff.Index, Line: "- " + lineDiff.Line})
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
