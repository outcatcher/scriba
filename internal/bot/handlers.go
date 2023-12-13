package bot

import (
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

type handlers struct {
	app schema.UseCases
	bot *telebot.Bot
}

// AddWorkflow adds workflow handling an endpoint.
func (h *handlers) AddWorkflow(endpoint string, w schema.Workflow) {
	w.WithUseCases(h.app)
	w.WithTelegramAPI(h.bot)

	h.bot.Handle(endpoint, w.EntryPoint(h.bot.Group()))
}
