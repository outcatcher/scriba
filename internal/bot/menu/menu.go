package menu

import (
	"errors"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

const (
	btnExit = "exit"
	btnBack = "back"

	textExit = "❌ выход"
	textBack = "🔙 назад"

	labelUserInfo = "user_info"
)

var (
	changeUserScoreButtons = [][]int32{
		{-1, +1},
		{-4, +4},
	}

	errMissingUser = errors.New("missing user info")
)

type selectedUserState struct {
	name       string
	telegramID int64
}

type userMenuState struct {
	app schema.UseCases
	bot *telebot.Bot
	grp *telebot.Group

	baseMsg     *telebot.Message // track /menu request
	baseMenuMsg *telebot.Message // track menu message to edit it in place

	currentUser  *selectedUserState
	currentLabel string
}

func newUserMenuState(app schema.UseCases, bot *telebot.Bot) *userMenuState {
	state := &userMenuState{
		app: app,
		bot: bot,
		grp: bot.Group(),
	}

	return state
}

func (u *userMenuState) back(c telebot.Context) error {
	switch u.currentLabel {
	case labelUserInfo:
		return u.selectPlayer(c)
	default:
		return nil
	}
}

func (u *userMenuState) exit(c telebot.Context) error {
	_ = c.Bot().Delete(u.baseMenuMsg)
	_ = c.Bot().Delete(u.baseMsg)

	return nil
}
