package entity

import (
	"github.com/cocoide/commitify-grpc-server/pkg/pb"
)

type LanguageType int

const (
	English LanguageType = iota + 1
	Japanese
)

type CodeFormatType int

const (
	NormalFormat = iota + 1
	PrefixFormat
	EmojiFormat
)

func ConvertPbLanguageToEntity(pbLanguage pb.LanguageType) LanguageType {
	return LanguageType(pbLanguage)
}

func ConvertPbCodeFormatToEntity(pbCodeFormat pb.CodeFormatType) CodeFormatType {
	return CodeFormatType(pbCodeFormat)
}
