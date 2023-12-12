package menu

import "gopkg.in/telebot.v3"

const (
	btnExit = "exit"
	btnBack = "back"

	textExit = "❌ выход"
	textBack = "🔙 назад"

	labelSelectPlayer = "select_player"
)

var changeUserScoreButtons = [][]int32{
	{-1, +1},
	{-4, +4},
}

func (u *userMenuState) back(c telebot.Context) error {
	switch u.previousLabel {
	case labelSelectPlayer:
		return u.selectPlayer(c)
	default:
		return nil
	}
}

func (u *userMenuState) exit(c telebot.Context) error {
	_ = u.api.Delete(u.baseMenuMsg)
	_ = u.api.Delete(u.baseMsg)

	return nil
}
