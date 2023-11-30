package bot

import (
	"context"
	"fmt"

	"github.com/outcatcher/scriba/internal/core/config"
	"gopkg.in/telebot.v3"
)

type botState struct {
	cfg config.BotConfig
}

// Start starts the bot.
func Start(ctx context.Context, cfg config.BotConfig) error {
	settings := telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{
			Timeout:        cfg.Poll.Interval,
			AllowedUpdates: []string{telebot.OnText}, // обрабатываем только сообщения
		},
		Verbose:   cfg.Verbose,
		ParseMode: telebot.ModeMarkdownV2,
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return fmt.Errorf("error creating new bot: %w", err)
	}

	bot.Use(autodelete(cfg))

	registerStartMenu(bot)

	go bot.Start()

	go func() {
		select {
		case <-ctx.Done():
			bot.Stop()
		}
	}()

	return nil
}
