package menu

import (
	"log/slog"
	"time"

	"gopkg.in/telebot.v3"
)

const deleteTimeout = 2 * time.Second

func errorReply(c telebot.Context, text string) {
	replyMsg, err := c.Bot().Reply(c.Message(), text)
	if err != nil {
		slog.Error("failed to reply with error", "error", err)

		return
	}

	time.AfterFunc(deleteTimeout, func() {
		err := c.Bot().Delete(replyMsg)
		if err != nil {
			slog.Error("failed to delete error message", "error", err)
		}
	})
}
