package schema

import "gopkg.in/telebot.v3"

// Workflow describes single command workflow.
type Workflow interface {
	// WithUseCases attaches Workflow to functionality.
	WithUseCases(useCases UseCases)
	// WithTelegramAPI attaches telegram API client.
	WithTelegramAPI(api TelegramAPI)
	// EntryPoint is a handler to start Workflow.
	EntryPoint(handler Handler) telebot.HandlerFunc
}
