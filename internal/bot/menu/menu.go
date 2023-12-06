package menu

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strconv"

	"github.com/outcatcher/scriba/internal/usecases"
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
	errOtherUser   = errors.New("clicked other user's menu")
)

type selectedUserState struct {
	name       string
	telegramID int64
}

type userMenuState struct {
	app *usecases.UseCases
	bot *telebot.Bot
	grp *telebot.Group

	baseMsg     *telebot.Message // track /menu request
	baseMenuMsg *telebot.Message // track menu message to edit it in place

	currentUser  *selectedUserState
	currentLabel string
}

func newUserMenuState(app *usecases.UseCases, bot *telebot.Bot) *userMenuState {
	state := &userMenuState{
		app: app,
		bot: bot,
		grp: bot.Group(),
	}

	state.grp.Use(state.restrictToAuthor)

	return state
}

func (u *userMenuState) restrictToAuthor(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		sender := c.Sender()

		if sender != nil && u.baseMsg != nil &&
			sender.ID != u.baseMsg.Sender.ID &&
			u.baseMenuMsg != nil { // bot own action
			slog.Error("user clicked another user menu", "base sender", u.baseMsg.Sender.ID, "this sender", sender.ID)

			return errOtherUser
		}

		return next(c)
	}
}

func (u *userMenuState) back(c telebot.Context) error {
	switch u.currentLabel {
	case labelUserInfo:
		return u.selectUser(c)
	default:
		return nil
	}
}

func (u *userMenuState) exit(c telebot.Context) error {
	_ = c.Bot().Delete(u.baseMenuMsg)
	_ = c.Bot().Delete(u.baseMsg)

	return nil
}

func (u *userMenuState) playersList(c telebot.Context, players []usecases.Player) []selectedUserState {
	result := make([]selectedUserState, 0, len(players))

	switch c.Chat().Type { //nolint:exhaustive
	case telebot.ChatPrivate:
		for _, player := range players {
			chat, err := u.bot.ChatByID(player.TelegramID)
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", c.Chat().ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, selectedUserState{
				name:       chat.FirstName,
				telegramID: player.TelegramID,
			})
		}
	case telebot.ChatGroup, telebot.ChatSuperGroup:
		for _, player := range players {
			member, err := u.bot.ChatMemberOf(c.Chat(), &telebot.User{ID: player.TelegramID})
			if err != nil {
				slog.Error("failed to get chat membership", "chat_id", c.Chat().ID, "user_id", player.TelegramID)

				continue
			}

			result = append(result, selectedUserState{
				name:       member.User.FirstName,
				telegramID: player.TelegramID,
			})
		}
	default:
		return nil
	}

	return result
}

func (u *userMenuState) selectUser(c telebot.Context) error {
	ctx := context.Background()
	msgText := "Выберите игрока"

	if u.baseMsg == nil {
		u.baseMsg = c.Message()
	}

	menu := &telebot.ReplyMarkup{
		OneTimeKeyboard: true,
		RemoveKeyboard:  true,
	}

	players, err := u.app.ListPlayers(ctx)
	if err != nil {
		return fmt.Errorf("error listing players: %w", err)
	}

	rows := make([]telebot.Row, 0, len(players)+1)

	for _, player := range u.playersList(c, players) {
		btn := menu.Data(player.name, strconv.FormatInt(player.telegramID, 10))

		selectedUser := selectedUserState{player.name, player.telegramID}

		u.grp.Handle(&btn, u.userDetails(ctx, selectedUser))

		rows = append(rows, menu.Row(btn))
	}

	exitBtn := menu.Data(textExit, btnExit)
	u.grp.Handle(&exitBtn, u.exit)

	rows = append(rows, telebot.Row{exitBtn})

	menu.Inline(rows...)

	if u.baseMenuMsg != nil {
		msg, err := c.Bot().Edit(u.baseMenuMsg, msgText, menu)
		if err != nil {
			return fmt.Errorf("error editing menu: %w", err)
		}

		u.baseMenuMsg = msg

		return nil
	}

	baseMenuMsg, err := c.Bot().Reply(c.Message(), msgText, menu)
	if err != nil {
		return fmt.Errorf("error sending reply: %w", err)
	}

	u.baseMenuMsg = baseMenuMsg

	return nil
}

func (u *userMenuState) changeScore(ctx context.Context, delta int32) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if u.currentUser == nil {
			return errMissingUser
		}

		// internal type is int16
		for delta > math.MaxInt16 {
			if err := u.app.UpdateCountByTelegramID(ctx, u.currentUser.telegramID, math.MaxInt16); err != nil {
				return fmt.Errorf("error updating player's count: %w", err)
			}

			delta -= math.MaxInt16
		}

		if err := u.app.UpdateCountByTelegramID(ctx, u.currentUser.telegramID, int16(delta)); err != nil {
			return fmt.Errorf("error updating player's count: %w", err)
		}

		return u.userDetails(ctx, *u.currentUser)(c)
	}
}

func uniqueByNumber(src int32) string {
	var prefix string
	if src < 0 {
		prefix = "minus"
	} else {
		prefix = "plus"
	}

	return fmt.Sprintf("%s%d", prefix, int32(math.Abs(float64(src))))
}

func (u *userMenuState) scoreButtonsToRows(
	ctx context.Context, menu *telebot.ReplyMarkup, src [][]int32,
) []telebot.Row {
	rows := make([]telebot.Row, len(src))

	for i, numRow := range changeUserScoreButtons {
		row := make(telebot.Row, len(numRow))

		for j, num := range numRow {
			btnValue := fmt.Sprintf("%+d", num)
			btn := menu.Data(btnValue, uniqueByNumber(num))

			u.grp.Handle(&btn, u.changeScore(ctx, num))

			row[j] = btn
		}

		rows[i] = row
	}

	return rows
}

func (u *userMenuState) userDetails(ctx context.Context, user selectedUserState) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		u.currentUser = &user
		u.currentLabel = labelUserInfo

		menu := &telebot.ReplyMarkup{
			OneTimeKeyboard: true,
			RemoveKeyboard:  true,
		}

		score, err := u.app.GetPlayerCountByTelegramID(ctx, user.telegramID)
		if err != nil {
			return fmt.Errorf("error getting player score: %w", err)
		}

		scoreStr := strconv.FormatInt(int64(score), 10)

		if score < 0 {
			scoreStr = "\\" + scoreStr // escape -
		}

		rows := u.scoreButtonsToRows(ctx, menu, changeUserScoreButtons)

		zeroBtn := menu.Data("Обнулить очки", "0")
		u.grp.Handle(&zeroBtn, u.changeScore(ctx, -score))

		rows = append(rows, telebot.Row{zeroBtn})

		exitBtn := menu.Data(textExit, btnExit)
		u.grp.Handle(&exitBtn, u.exit)

		backBtn := menu.Data(textBack, btnBack)
		u.grp.Handle(&backBtn, u.back)

		// some other data here
		rows = append(rows, telebot.Row{backBtn, exitBtn})

		menu.Inline(rows...)

		messageText := fmt.Sprintf("Количество баллов у игрока %s: %s", user.name, scoreStr)

		msg, err := c.Bot().Edit(u.baseMenuMsg, messageText, menu)
		if err != nil {
			return fmt.Errorf("error editing menu: %w", err)
		}

		u.baseMenuMsg = msg

		return nil
	}
}
