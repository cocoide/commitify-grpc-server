package service

import (
	"github.com/cocoide/commitify-grpc-server/internal/domain/entity"
)

//go:generate mockgen -source=third_party.go -destination=../../pkg/mockmock/deepl.go
type LangService interface {
	TranslateTexts(texts []string, into entity.LanguageType) ([]string, error)
	TranslateTextsIntoJapanese(texts []string) ([]string, error)
}

//go:generate mockgen -source=third_party.go -destination=../../pkg/mock/openai.go
type NLPService interface {
	GetAnswerFromPrompt(prompt string) (string, error)
}
