package register

import (
	"fmt"
	"log/slog"

	"github.com/outcatcher/scriba/internal/bot/common"
	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

const registerUnique = "register"

type state struct {
	app     schema.UseCases    // application logic handlers
	api     schema.TelegramAPI // interaction with telegram API
	handler schema.Handler     // handler to handle subsequent commands (can be Bot or Group)
}

func (s *state) handleStartMenu(c telebot.Context) error {
	response := &telebot.ReplyMarkup{}

	menu := &telebot.ReplyMarkup{}
	registerBtn := menu.Data("Хочу играть!", registerUnique)

	response.Inline(menu.Row(registerBtn))

	_, err := s.api.Send(
		c.Chat(),
		"Привет\\!\nЕсли хочешь, чтобы для тебя считались баллы, нажми на кнопку\\.",
		response,
	)
	if err != nil {
		return fmt.Errorf("failed to handleStart: %w", err)
	}

	s.handler.Handle(registerBtn, s.handleRegisterButton)

	return nil
}

func (s *state) handleRegisterButton(c telebot.Context) error {
	ctx := common.GetContextFromContext(c)

	sender := c.Sender()

	replyText := fmt.Sprintf("%s теперь игрок\\!", sender.FirstName)

	err := s.app.RegisterWithTelegram(ctx, sender.ID)
	if err != nil {
		slog.Error("failed to register user", "error", err)

		_, err := s.api.Reply(c.Message(), "Не смогли зарегистрировать пользователя :\\(")
		if err != nil {
			return fmt.Errorf("failed to register user reply error: %w", err)
		}

		return nil
	}

	_, err = s.api.Reply(c.Message(), replyText)
	if err != nil {
		return fmt.Errorf("failed to handle handleStart reply: %w", err)
	}

	return nil
}
