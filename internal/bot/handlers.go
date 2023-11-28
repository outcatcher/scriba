package bot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

const (
	parentResponse = "parent"
	childResponse  = "child"
)

var (
	startMenu          = &telebot.ReplyMarkup{}
	startMenuParentBtn = startMenu.Data("Я родитель 👨‍🦳👩‍🦳", parentResponse)
	startMenuChildBtn  = startMenu.Data("Я ребёнок 👦👧", childResponse)
)

func registerStartMenu(bot *telebot.Bot) {
	bot.Handle("/start", start)
	bot.Handle(&startMenuParentBtn, startReply(parentResponse))
	bot.Handle(&startMenuChildBtn, startReply(childResponse))
}

func start(ctx telebot.Context) error {
	response := &telebot.ReplyMarkup{}

	response.Inline(startMenu.Row(startMenuParentBtn, startMenuChildBtn))

	err := ctx.Send("Привет\\!\nОтметь, кем ты являешься в нашей игре\\.", response)
	if err != nil {
		return fmt.Errorf("failed to say start: %w", err)
	}

	return nil
}

func localizeStartReply(role string) string {
	switch role {
	case childResponse:
		return "ребёнком"
	case parentResponse:
		return "родителем"
	default:
		return ""
	}
}

func startReply(role string) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		replyText := fmt.Sprintf("%s теперь считается *%s*", ctx.Sender().FirstName, localizeStartReply(role))

		err := ctx.Reply(replyText)
		if err != nil {
			return fmt.Errorf("failed to handle start reply: %w", err)
		}

		return nil
	}
}
