package entity

import (
	pb "github.com/cocoide/commitify-grpc-server/pkg/grpc"
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

func (l LanguageType) ConvertToDeeplType() string {
	switch l {
	case Japanese:
		return "JA"
	case English:
		return "EN"
	default:
		return ""
	}
}
