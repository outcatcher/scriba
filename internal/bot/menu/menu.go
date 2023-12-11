package menu

import (
	"errors"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

var errMissingUser = errors.New("missing user info")

type userInfo struct {
	name       string
	telegramID int64
}

type userMenuState struct {
	app schema.UseCases // application logic handlers

	handler schema.Handler // handler to handle subsequent commands (can be Bot or Group)

	baseMsg     telebot.Editable // track /menu request
	baseMenuMsg telebot.Editable // track menu message to edit it in place

	selectedUser *userInfo

	currentLabel string
}
