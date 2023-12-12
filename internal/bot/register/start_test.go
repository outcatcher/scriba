package register

import (
	"context"
	"errors"
	"testing"

	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/telebot.v3"
)

var (
	expectedChat   = &telebot.Chat{Type: telebot.ChatGroup, ID: 1}
	expectedSender = &telebot.User{ID: 2, FirstName: "tester"}
)

type workflowSuite struct {
	suite.Suite

	tgMock      *mocks.MockTelegramAPI
	appMock     *mocks.MockUseCases
	handlerMock *mocks.MockHandler
	contextMock *mocks.MockContext

	workflow *Workflow
}

// suite with such setup/teardown can't have parallel tests

func (s *workflowSuite) SetupTest() {
	s.tgMock = new(mocks.MockTelegramAPI)
	s.appMock = new(mocks.MockUseCases)
	s.handlerMock = new(mocks.MockHandler)
	s.contextMock = new(mocks.MockContext)

	s.workflow = new(Workflow)

	s.workflow.WithUseCases(s.appMock)
	s.workflow.WithTelegramAPI(s.tgMock)
}

func (s *workflowSuite) TearDownTest() {
	t := s.T()

	s.tgMock.AssertExpectations(t)
	s.appMock.AssertExpectations(t)
	s.handlerMock.AssertExpectations(t)
	s.contextMock.AssertExpectations(t)
}

func (s *workflowSuite) TestStartWorkflow() { //nolint:funlen  // the simplest way used
	t := s.T()

	start := s.workflow.EntryPoint(s.handlerMock)

	ctx := context.Background()

	s.contextMock.On("Get", "timeout-context").Return(ctx)
	s.contextMock.On("Chat").Return(expectedChat)
	s.contextMock.On("Sender").Return(expectedSender)

	expectedMsg := &telebot.Message{Chat: expectedChat}
	s.tgMock.
		On("Send", expectedChat, mock.AnythingOfType("string"), mock.AnythingOfType("*telebot.ReplyMarkup")).
		Run(func(args mock.Arguments) {
			markup, ok := args.Get(2).(*telebot.ReplyMarkup)
			require.True(t, ok)

			expectedMsg.ReplyMarkup = markup
		}).
		Return(expectedMsg, nil)

	var clickRegister telebot.HandlerFunc

	s.handlerMock.
		On("Handle", mock.Anything, mock.AnythingOfType("telebot.HandlerFunc")).
		Run(func(args mock.Arguments) {
			fn, ok := args.Get(1).(telebot.HandlerFunc)
			require.True(t, ok)

			clickRegister = fn
		}).
		Return(nil)

	// go through register workflow
	require.NoError(t, start(s.contextMock))

	cases := []struct {
		name    string
		err     error
		inReply string
	}{
		{"ok", nil, expectedSender.FirstName},
		{"fail", errors.New("TEST ERROR"), "Не смогли"}, //nolint:goerr113
	}

	for i := range cases {
		data := cases[i]

		t.Run(data.name, func(t *testing.T) {
			s.appMock.On("RegisterWithTelegram", ctx, expectedSender.ID).Return(data.err).Once()
			s.contextMock.On("Message").Return(expectedMsg).Once()

			expectedReply := &telebot.Message{ReplyTo: expectedMsg}
			s.tgMock.
				On("Reply", expectedMsg, mock.AnythingOfType("string")).
				Run(func(args mock.Arguments) {
					expectedReply.Text = args.String(1)

					require.Contains(t, args.String(1), data.inReply)
				}).
				Return(expectedReply, nil).Once()

			require.NoError(t, clickRegister(s.contextMock))
		})
	}
}

func TestWorkflow(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(workflowSuite))
}
