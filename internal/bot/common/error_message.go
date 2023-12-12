package common

import (
	"log/slog"
	"time"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

const deleteTimeout = 2 * time.Second

// ErrorReply - reply with internal error.
func ErrorReply(replier schema.TelegramAPI, message *telebot.Message, text string) {
	replyMsg, err := replier.Reply(message, text)
	if err != nil {
		slog.Error("failed to reply with error", "error", err)

		return
	}

	time.AfterFunc(deleteTimeout, func() {
		err := replier.Delete(replyMsg)
		if err != nil {
			slog.Error("failed to delete error message", "error", err)
		}
	})
}
