package menu

import (
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

// UserMenuWorkflow is workflow for /menu command.
type UserMenuWorkflow struct {
	app schema.UseCases
}

// WithUseCases attaches workflow to functionality.
func (uw *UserMenuWorkflow) WithUseCases(useCases schema.UseCases) {
	uw.app = useCases
}

// EntryPoint is a handler to start /menu workflow.
func (uw *UserMenuWorkflow) EntryPoint(handler schema.Handler) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if handler == nil {
			handler = c.Bot()
		}

		state := &userMenuState{
			app:     uw.app,
			handler: handler,
		}

		return state.selectPlayer(c)
	}
}
