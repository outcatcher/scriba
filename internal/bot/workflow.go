package bot

import (
	"github.com/outcatcher/scriba/internal/usecases"
	"gopkg.in/telebot.v3"
)

type workflow interface {
	// WithUseCases attaches workflow to functionality.
	WithUseCases(useCases *usecases.UseCases)
	// Start is a handler to start workflow.
	Start(bot *telebot.Bot) telebot.HandlerFunc
}

func (h *handlers) addWorkflow(bot *telebot.Bot, endpoint string, w workflow) {
	w.WithUseCases(h.app)

	bot.Handle(endpoint, w.Start(bot))
}
