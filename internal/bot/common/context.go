package common

import (
	"context"
	"time"

	"gopkg.in/telebot.v3"
)

const contextKeyContext = "timeout-context"

// WithTimeoutContext creates context from base one and passes it into telebot context.
func WithTimeoutContext(baseCtx context.Context, timeout time.Duration) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			ctx, cancel := context.WithTimeout(baseCtx, timeout)
			defer cancel()

			c.Set(contextKeyContext, ctx)

			return next(c)
		}
	}
}

// GetContextFromContext returns context stored in context.
func GetContextFromContext(c telebot.Context) context.Context {
	val := c.Get(contextKeyContext)

	if ctx, ok := val.(context.Context); ok {
		return ctx
	}

	return context.Background() //nolint:forbidigo  // fallback context
}
