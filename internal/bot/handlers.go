package bot

import (
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

type handlers struct {
	app schema.UseCases
	bot *telebot.Bot
}

type workflow interface {
	// WithUseCases attaches workflow to functionality.
	WithUseCases(useCases schema.UseCases)
	// WithTelegramAPI attaches telegram API client.
	WithTelegramAPI(api schema.TelegramAPI)
	// EntryPoint is a handler to start workflow.
	EntryPoint(handler schema.Handler) telebot.HandlerFunc
}

func (h *handlers) addWorkflow(endpoint string, w workflow) {
	w.WithUseCases(h.app)
	w.WithTelegramAPI(h.bot)

	h.bot.Handle(endpoint, w.EntryPoint(h.bot.Group()))
}
