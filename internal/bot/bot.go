package bot

import (
	"context"
	"fmt"

	"github.com/outcatcher/scriba/internal/core/config"
	"gopkg.in/telebot.v3"
)

// Start starts the bot.
func Start(ctx context.Context, cfg *config.Configuration) error {
	settings := telebot.Settings{
		Token: cfg.Bot.Token,
		Poller: &telebot.LongPoller{
			Timeout:        cfg.Bot.Poll.Interval,
			AllowedUpdates: []string{telebot.OnText}, // обрабатываем только сообщения
		},
		Verbose:   cfg.Bot.Verbose,
		ParseMode: telebot.ModeMarkdownV2,
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return fmt.Errorf("error creating new bot: %w", err)
	}

	bot.Use(autodelete(cfg.Bot))

	hdl, err := newHandlers(cfg)
	if err != nil {
		if err != nil {
			return fmt.Errorf("error creating new bot: %w", err)
		}
	}

	hdl.registerHandlers(bot) //nolint:contextcheck

	go bot.Start()

	go func() {
		<-ctx.Done()
		bot.Stop()
	}()

	return nil
}
