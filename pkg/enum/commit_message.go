package enum

import "github.com/cocoide/commitify-grpc-server/pkg/pb"

type Language int

const (
	Japanese Language = iota
	English
)

type CodeFormat int

const (
	EmojiFormat CodeFormat = iota
	PrefixFormat
	NormalFormat
)

func ConvertPbLanguage(pbLanguage pb.LanguageType) Language {
	switch pbLanguage {
	case pb.LanguageType_JAPANESE:
		return Japanese
	case pb.LanguageType_ENGLISH:
		return English
	default:
		return Japanese // default
	}
}

func ConvertPbCodeFormat(pbCodeFormat pb.CodeFormatType) CodeFormat {
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
