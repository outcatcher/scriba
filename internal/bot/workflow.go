package bot

import (
	"context"

	"github.com/outcatcher/scriba/internal/entities"
	"gopkg.in/telebot.v3"
)

type UseCases interface {
	UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error
	GetPlayerCountByTelegramID(ctx context.Context, telegramID int64) (int32, error)
	RegisterWithTelegram(ctx context.Context, telegramID int64) error
	ListPlayers(ctx context.Context) ([]entities.Player, error)
}

type workflow interface {
	// WithUseCases attaches workflow to functionality.
	WithUseCases(useCases UseCases)
	// Start is a handler to start workflow.
	Start(bot *telebot.Bot) telebot.HandlerFunc
}

func (h *handlers) addWorkflow(bot *telebot.Bot, endpoint string, w workflow) {
	w.WithUseCases(h.app)

	bot.Handle(endpoint, w.Start(bot))
}
