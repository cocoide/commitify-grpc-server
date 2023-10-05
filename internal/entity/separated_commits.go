package entity

import pb "github.com/cocoide/commitify-grpc-server/pkg/grpc"

type SeparatedCommitMessage struct {
	Messages   []string
	Filename   string
	ChangeType ChangeType
}

type FileChange struct {
	CodeDiff   CodeDiff
	Filename   string
	ChangeType ChangeType
}

type LineDiff struct {
	Index int32
	Line  string
}

type CodeDiff struct {
	Added   []LineDiff
	Deleted []LineDiff
}

type ChangeType int

const (
	CREATE ChangeType = iota + 1
	UPDATE
	DELETE
)

func (c ChangeType) String() string {
	switch c {
	case CREATE:
		return "CREATE"
	case UPDATE:
		return "UPDATE"
	case DELETE:
		return "DELETE"
	default:
		return "UNKOWN"
	}
}

func (c *SeparatedCommitMessage) ConvertToPbCommitMessages() *pb.SeparatedCommitMessages {
	return &pb.SeparatedCommitMessages{
		Messages:   c.Messages,
		Filename:   c.Filename,
		ChangeType: pb.ChangeType(c.ChangeType),
	}
}

func ConvertPbToFileChange(pbFileChange *pb.FileChange) FileChange {
	fileChange := FileChange{
		Filename:   pbFileChange.Filename,
		ChangeType: ChangeType(pbFileChange.ChangeType),
	}

	for _, pbLineDiff := range pbFileChange.CodeDiff.Added {
		lineDiff := LineDiff{
			Index: pbLineDiff.Index,
			Line:  pbLineDiff.Line,
		}
		fileChange.CodeDiff.Added = append(fileChange.CodeDiff.Added, lineDiff)
	}

	for _, pbLineDiff := range pbFileChange.CodeDiff.Deleted {
		lineDiff := LineDiff{
			Index: pbLineDiff.Index,
			Line:  pbLineDiff.Line,
		}
		fileChange.CodeDiff.Deleted = append(fileChange.CodeDiff.Deleted, lineDiff)
	}
	return fileChange
}
