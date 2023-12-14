package menu

import (
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

// Workflow is workflow for /menu command.
type Workflow struct {
	app schema.UseCases
	api schema.TelegramAPI
}

// WithUseCases attaches workflow to functionality.
func (w *Workflow) WithUseCases(useCases schema.UseCases) {
	w.app = useCases
}

// WithTelegramAPI attaches telegram API client.
func (w *Workflow) WithTelegramAPI(api schema.TelegramAPI) {
	w.api = api
}

// EntryPoint is a handler to start /menu workflow.
func (w *Workflow) EntryPoint(handler schema.Handler) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		state := &userMenuState{
			app:     w.app,
			api:     w.api,
			handler: handler,
		}

		return state.selectPlayer(c)
	}
}
