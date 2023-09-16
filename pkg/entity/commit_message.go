package entity

import "github.com/cocoide/commitify-grpc-server/pkg/pb"

type LanguageType int

const (
	Japanese LanguageType = iota
	English
	OtherLanguage
)

type CodeFormatType int

const (
	EmojiFormat CodeFormatType = iota
	PrefixFormat
	NormalFormat
)

func ConvertPbLanguageToEntity(pbLanguage pb.LanguageType) LanguageType {
	switch pbLanguage {
	case pb.LanguageType_JAPANESE:
		return Japanese
	case pb.LanguageType_ENGLISH:
		return English
	default:
		return Japanese // default
	}
}

func ConvertPbCodeFormatToEntity(pbCodeFormat pb.CodeFormatType) CodeFormatType {
	switch pbCodeFormat {
	case pb.CodeFormatType_EMOJI:
		return EmojiFormat
	case pb.CodeFormatType_PREFIX:
		return PrefixFormat
	case pb.CodeFormatType_NORMAL:
		return NormalFormat
	default:
		return NormalFormat // default
	}
}
