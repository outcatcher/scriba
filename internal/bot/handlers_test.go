package bot

import (
	"testing"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gopkg.in/telebot.v3"
)

func TestHandlers_AddWorkflow(t *testing.T) {
	t.Parallel()

	ucMock := mocks.NewMockUseCases(t)
	wflw := mocks.NewMockWorkflow(t)

	wflw.On("WithUseCases", ucMock).Return()

	bot, err := telebot.NewBot(telebot.Settings{Offline: true})
	require.NoError(t, err)

	handls := handlers{
		app: ucMock,
		bot: bot,
	}

	wflw.On("WithTelegramAPI", bot).Return()
	wflw.On("EntryPoint", mock.Anything).Run(func(args mock.Arguments) {
		require.Implements(t, new(schema.Handler), args.Get(0))
	}).Return(emptyHandler())

	handls.AddWorkflow("test", wflw)
}

func emptyHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error { return nil }
}
