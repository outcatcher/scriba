package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/outcatcher/scriba/internal/bot/common"
	"github.com/outcatcher/scriba/internal/bot/schema"
	"github.com/outcatcher/scriba/internal/bot/workflows/menu"
	"github.com/outcatcher/scriba/internal/bot/workflows/register"
	"github.com/outcatcher/scriba/internal/core/config"
	"gopkg.in/telebot.v3"
)

const handlerTimeout = 10 * time.Second

// Start starts the bot.
func Start(ctx context.Context, botConfig config.BotConfig, app schema.UseCases) error {
	settings := telebot.Settings{
		Token: botConfig.Token,
		Poller: &telebot.LongPoller{
			Timeout:        botConfig.Poll.Interval,
			AllowedUpdates: []string{telebot.OnText}, // handling text messages only
		},
		Verbose:   botConfig.Verbose,
		ParseMode: telebot.ModeMarkdownV2,
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return fmt.Errorf("error creating new bot: %w", err)
	}

	bot.Use(
		common.LogCommands(),
		common.WithTimeoutContext(ctx, handlerTimeout),
	)

	if botConfig.Messages != nil {
		bot.Use(common.AutodeleteMiddleware(botConfig.Messages.Lifetime))
	}

	hdl := &handlers{
		app: app,
		bot: bot,
	}
	hdl.AddWorkflow("/start", new(register.Workflow))
	hdl.AddWorkflow("/menu", new(menu.Workflow))

	go func() {
		<-ctx.Done()
		bot.Stop()
	}()

	bot.Start()

	return nil
}
