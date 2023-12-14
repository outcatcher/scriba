package menu

import (
	"errors"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

var errMissingUser = errors.New("user is not selected")

type telegramUserInfo struct {
	name       string
	telegramID int64
}

type userMenuState struct {
	app     schema.UseCases    // application logic handlers
	api     schema.TelegramAPI // telegram bot API client
	handler schema.Handler     // handler to handle subsequent commands (can be Bot or Group)

	baseMsg     telebot.Editable // track /menu request
	baseMenuMsg telebot.Editable // track menu message to edit it in place

	selectedUser *telegramUserInfo // selected user, if any

	previousLabel string // label of previous menu
}
