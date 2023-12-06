package bot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/outcatcher/scriba/internal/bot/menu"
	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/usecases"
	"gopkg.in/telebot.v3"
)

var (
	startMenu         = &telebot.ReplyMarkup{}
	startMenuChildBtn = startMenu.Data("Хочу играть!", "register")
)

type handlers struct {
	app *usecases.UseCases
}

func newHandlers(cfg *config.Configuration) (*handlers, error) {
	app, err := initApp(cfg)
	if err != nil {
		return nil, fmt.Errorf("error intializing app: %w", err)
	}

	return &handlers{
		app: app,
	}, nil
}

func (h *handlers) registerHandlers(bot *telebot.Bot) {
	bot.Handle("/start", h.handleStart)
	bot.Handle(&startMenuChildBtn, h.handleStartReply)

	h.addWorkflow(bot, "/menu", new(menu.UserMenuWorkflow))
}

func (*handlers) handleStart(c telebot.Context) error {
	response := &telebot.ReplyMarkup{}

	response.Inline(startMenu.Row(startMenuChildBtn))

	err := c.Send("Привет\\!\nЕсли хочешь, чтобы для тебя считались баллы, нажми на кнопку\\.", response)
	if err != nil {
		return fmt.Errorf("failed to handleStart: %w", err)
	}

	return nil
}

func (h *handlers) handleStartReply(c telebot.Context) error {
	sender := c.Sender()

	replyText := fmt.Sprintf("%s теперь игрок\\!", sender.FirstName)

	err := h.app.RegisterWithTelegram(context.Background(), sender.ID)
	if err != nil {
		slog.Error("failed to register user: %w", "error", err)

		err := c.Reply("Не смогли зарегистрировать пользователя :\\(")
		if err != nil {
			return fmt.Errorf("failed to register user reply error: %w", err)
		}

		return nil
	}

	if err := c.Reply(replyText); err != nil {
		return fmt.Errorf("failed to handle handleStart reply: %w", err)
	}

	return nil
}
