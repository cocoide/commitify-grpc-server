package usecase

import (
	"fmt"
	"github.com/cocoide/commitify-grpc-server/internal/entity"
	"github.com/cocoide/commitify-grpc-server/internal/gateway"
	"regexp"
	"strings"

	"github.com/cocoide/commitify-grpc-server/utils"
)

const (
	NormalMessagePrompt = "Generate up to 5 commit messages for [%s]. Each message should be separated by only space"
	EmojiMessagePrompt  = "Select the appropriate emoji for each commit messages [%v] from the following emoji map {emoji: meaning, %s} and return the emojis separated by commas"
	PrefixMessagePrompt = "Select the appropriate prefix for each commit messages [%v] from the following prefix map {prefix: meaning, %s} and return the prefixs separated by commas."
)

var messagesRegex = regexp.MustCompile(`^(\d.\s+)|^(-\s+)|^(\s+)`)

type CommitMessageUsecase struct {
	og gateway.OpenAIGateway
	dg gateway.DeeplAPIGateway
}

func NewCommitMessageUsecaes(og gateway.OpenAIGateway, dg gateway.DeeplAPIGateway) *CommitMessageUsecase {
	return &CommitMessageUsecase{og: og, dg: dg}
}

func generateEnglishMessage(og gateway.OpenAIGateway, code string) ([]string, error) {
	maxCodeLength := 200
	if len(code) > maxCodeLength {
		code = code[:maxCodeLength]
	}
	prompt := fmt.Sprintf(NormalMessagePrompt, code)
	result, err := og.GetAnswerFromPrompt(prompt)
	if err != nil {
		return nil, err
	}
	messages := strings.Split(result, "\n")
	messages = utils.RemoveFromArrayByRegex(messages, messagesRegex)
	return messages, nil
}

func (u *CommitMessageUsecase) GenerateNormalMessage(code string, language entity.LanguageType) ([]string, error) {
	messages, err := generateEnglishMessage(u.og, code)
	if err != nil {
		return nil, err
	}
	if language == entity.Japanese {
		messages, err = u.dg.TranslateTextsIntoJapanese(messages)
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
}

func (u *CommitMessageUsecase) GenerateEmojiMessage(code string, language entity.LanguageType) ([]string, error) {
	messages, err := generateEnglishMessage(u.og, code)
	if err != nil {
		return []string{}, err
	}
	emojiMap := make(map[string]string, 6)
	emojiMap["🐛"] = "bugfix"
	emojiMap["🎉"] = "release"
	emojiMap["✨"] = "update"
	emojiMap["📄"] = "documentation"
	emojiMap["🔓"] = "security"
	emojiMap["⚡️"] = "performance"
	emojiMap["🗑️"] = "delete"
	prompt := fmt.Sprintf(EmojiMessagePrompt, messages, emojiMap)
	emojiLine, err := u.og.GetAnswerFromPrompt(prompt)
	if err != nil {
		return nil, err
	}
	emojis := strings.Split(strings.ReplaceAll(emojiLine, " ", ""), ",")

	if language == entity.Japanese {
		messages, err = u.dg.TranslateTextsIntoJapanese(messages)
		if err != nil {
			return nil, err
		}
	}
	for i, emoji := range emojis {
		if _, ok := emojiMap[emoji]; !ok {
			emoji = "✨"
		}
		messages[i] = emoji + " " + messages[i]
	}
	return messages, nil
}

func (u *CommitMessageUsecase) GeneratePrefixMessage(code string, language entity.LanguageType) ([]string, error) {
	messages, err := generateEnglishMessage(u.og, code)
	prefixMap := make(map[string]string, 6)
	prefixMap["feat"] = "feature"
	prefixMap["fix"] = "bugfix"
	prefixMap["docs"] = "document"
	prefixMap["style"] = "format"
	prefixMap["perf"] = "performance"
	prefixMap["chore"] = "unimportant"
	prompt := fmt.Sprintf(PrefixMessagePrompt, messages, prefixMap)
	prefixLine, err := u.og.GetAnswerFromPrompt(prompt)
	if err != nil {
		return nil, err
	}
	prefixs := strings.Split(strings.ReplaceAll(prefixLine, " ", ""), ",")
	if err != nil {
		return nil, err
	}
	for i, prefix := range prefixs {
		if _, ok := prefixMap[prefix]; !ok {
			prefix = "feat"
		}
		messages[i] = prefix + ": " + messages[i]
	}
	if language == entity.Japanese {
		messages, err = u.dg.TranslateTextsIntoJapanese(messages)
		if err != nil {

		}
	}
	return messages, nil
}
