package common

import (
	"log/slog"
	"time"

	"github.com/outcatcher/scriba/internal/bot/schema"
	"gopkg.in/telebot.v3"
)

var deleteTimeout = 2 * time.Second // made a var to simplify tests  todo: move to configuration

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
