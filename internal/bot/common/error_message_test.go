package common

import (
	"testing"
	"time"

	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/mock"
	"gopkg.in/telebot.v3"
)

func TestErrorReply(t *testing.T) {
	t.Parallel()

	deleteTimeout = time.Millisecond

	mockTg := mocks.NewMockTelegramAPI(t)

	message := &telebot.Message{ID: 1}
	replyMsg := &telebot.Message{ID: 2, ReplyTo: message}

	expectedText := "hi!"

	mockTg.On("Reply", message, mock.AnythingOfType("string")).Run(func(args mock.Arguments) {
		text := args.String(1)

		replyMsg.Text = text
	}).Return(replyMsg, nil)
	mockTg.On("Delete", replyMsg).Return(nil)

	ErrorReply(mockTg, message, expectedText)

	time.Sleep(deleteTimeout * 2)
}
