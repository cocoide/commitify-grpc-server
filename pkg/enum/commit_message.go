package enum

import "github.com/cocoide/commitify-grpc-server/pkg/pb"

const (
	JAPANESE = pb.LanguageType_JAPANESE
	ENGLISH  = pb.LanguageType_ENGLISH
)
const (
	EMOJI_FORMAT  = pb.CodeFormatType_EMOJI
	PREFIX_FORMAT = pb.CodeFormatType_PREFIX
	NORMAL_FORMAT = pb.CodeFormatType_NORMAL
)
