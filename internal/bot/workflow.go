package bot

import (
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

type workflow interface {
	// WithUseCases attaches workflow to functionality.
	WithUseCases(useCases schema.UseCases)
	// Start is a handler to start workflow.
	Start(bot *telebot.Bot) telebot.HandlerFunc
}

func (h *handlers) addWorkflow(bot *telebot.Bot, endpoint string, w workflow) {
	w.WithUseCases(h.app)

	bot.Handle(endpoint, w.Start(bot))
}
