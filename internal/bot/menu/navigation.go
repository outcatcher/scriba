package menu

import "gopkg.in/telebot.v3"

const (
	btnExit = "exit"
	btnBack = "back"

	textExit = "❌ выход"
	textBack = "🔙 назад"

	labelUserInfo = "user_info"
)

var changeUserScoreButtons = [][]int32{
	{-1, +1},
	{-4, +4},
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
