package menu

import (
	"github.com/outcatcher/scriba/internal/bot"
	"gopkg.in/telebot.v3"
)

// UserMenuWorkflow is workflow for /menu command.
type UserMenuWorkflow struct {
	app bot.UseCases
}

// WithUseCases attaches workflow to functionality.
func (uw *UserMenuWorkflow) WithUseCases(useCases bot.UseCases) {
	uw.app = useCases
}

// Start is a handler to start /menu workflow.
func (uw *UserMenuWorkflow) Start(bot *telebot.Bot) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		state := newUserMenuState(uw.app, bot)

		return state.selectPlayer(c)
	}
}
