package bot

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/outcatcher/scriba/internal/core/config"
	"github.com/outcatcher/scriba/internal/usecases"
	"gopkg.in/telebot.v3"
)

const (
	register = "register"

	directionUp   = "up"
	directionDown = "down"

	deleteTimeout = time.Minute
)

var (
	startMenu         = &telebot.ReplyMarkup{}
	startMenuChildBtn = startMenu.Data("Хочу играть!", register)
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

	bot.Handle("/menu", h.handleMenu)

	bot.Handle("/up", h.handleCountChange(directionUp))
	bot.Handle("/down", h.handleCountChange(directionDown))

	bot.Handle("/stat", h.handleStat)
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

func directionMultiplier(direction string) int16 {
	switch direction {
	case directionUp:
		return 1
	case directionDown:
		return -1
	default:
		return 0
	}
}

func (h *handlers) handleCountChange(direction string) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		replyTo := c.Message().ReplyTo
		if replyTo == nil {
			err := c.Reply("Оценка должна быть ответом на сообщение")
			if err != nil {
				return fmt.Errorf("failed to handle handleCountChange error: %w", err)
			}

			return nil
		}

		recipientID := replyTo.Sender.ID

		args := c.Args()
		if len(args) < 1 {
			err := c.Reply("Недостаточно аргументов")
			if err != nil {
				return fmt.Errorf("failed to handle handleCountChange error: %w", err)
			}

			return nil
		}

		delta, err := strconv.ParseInt(args[0], 10, 16)
		if err != nil {
			slog.Error("error handling count change", "error", err)

			err := c.Reply("Оценка должна быть числом")
			if err != nil {
				return fmt.Errorf("failed to handle handleCountChange error: %w", err)
			}

			return nil
		}

		err = h.app.UpdateCountByTelegramID(
			context.Background(), recipientID, directionMultiplier(direction)*int16(delta),
		)
		if err != nil {
			slog.Error("error handling count change", "error", err)

			err := c.Reply("Не смогли обновить счёт :\\(")
			if err != nil {
				return fmt.Errorf("failed to handle handleCountChange error: %w", err)
			}

			return nil
		}

		return nil
	}
}

func (h *handlers) handleStat(c telebot.Context) error {
	replyTo := c.Message().ReplyTo
	if replyTo == nil {
		err := c.Reply("Запрос статистики должен быть ответом на сообщение")
		if err != nil {
			return fmt.Errorf("failed to handle handleStat error: %w", err)
		}

		return nil
	}

	count, err := h.app.GetPlayerCountByTelegramID(context.Background(), replyTo.Sender.ID)
	if err != nil {
		slog.Error("error handling get count", "error", err)

		err := c.Reply("Не смогли получить счёт :\\(")
		if err != nil {
			return fmt.Errorf("failed to handle handleStat error: %w", err)
		}

		return nil
	}

	countStr := strconv.FormatInt(int64(count), 10)

	if count < 0 {
		countStr = "\\" + countStr // escape -
	}

	err = c.Reply(fmt.Sprintf("Количество баллов у игрока %s: %s", replyTo.Sender.FirstName, countStr))
	if err != nil {
		return fmt.Errorf("failed to reply handleStat: %w", err)
	}

	return nil
}

func (h *handlers) handleMenu(c telebot.Context) error {
	menu, err := h.selectUserMenu(c)
	if err != nil {
		slog.Error("error handling base menu", "error", err)

		c.DeleteAfter(deleteTimeout)

		err := c.Reply("Не смогли отобразить меню :\\(")
		if err != nil {
			return fmt.Errorf("failed to handle handleStat error: %w", err)
		}

		return nil
	}

	if err := c.Reply(menu); err != nil {
		return fmt.Errorf("failed to reply base menu: %w", err)
	}

	return nil
}
