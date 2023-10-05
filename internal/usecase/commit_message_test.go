package usecase_test

import (
	"fmt"
	"github.com/cocoide/commitify-grpc-server/internal/domain/entity"
	"github.com/cocoide/commitify-grpc-server/internal/usecase"
	mock_gateway "github.com/cocoide/commitify-grpc-server/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommitMessageUsecaseTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	openai  *mock_gateway.MockOpenAIGateway
	deepl   *mock_gateway.MockDeeplAPIGateway
	usecase *usecase.CommitMessageUsecase
}

func (u *CommitMessageUsecaseTestSuite) SetupTest() {
	u.ctrl = gomock.NewController(u.T())
	u.openai = mock_gateway.NewMockOpenAIGateway(u.ctrl)
	u.deepl = mock_gateway.NewMockDeeplAPIGateway(u.ctrl)
	u.usecase = usecase.NewCommitMessageUsecaes(u.openai, u.deepl)
}

func (u *CommitMessageUsecaseTestSuite) TearDonwTest() {
	u.ctrl.Finish()
}

func (u *CommitMessageUsecaseTestSuite) TestGenerateNormalMessage() {
	type testcase struct {
		name     string
		language entity.LanguageType
		err      error
		runMock  func()
	}
	code := "Add funcA and update its READ.md"
	//englishCommitMessages := []string{"feat: add funcA", "docs: update READ.md"}
	tests := []testcase{
		{"OK", entity.English, nil,
			func() {
				u.openai.EXPECT().
					GetAnswerFromPrompt(fmt.Sprintf(usecase.NormalMessagePrompt, code)).AnyTimes()
				u.deepl.EXPECT().
					TranslateTextsIntoJapanese(gomock.Any()).AnyTimes()
			},
		},
	}

	for _, tt := range tests {
		u.T().Run(tt.name, func(t *testing.T) {
			tt.runMock()
			_, err := u.usecase.GenerateNormalMessage(code, tt.language)
			if err != nil {
				u.Error(err)
			}
		})
	}
}

func Test_CommitMessageUseCase(t *testing.T) {
	suite.Run(t, new(CommitMessageUsecaseTestSuite))
}
