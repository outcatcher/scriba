package common

import (
	"context"
	"testing"
	"time"

	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gopkg.in/telebot.v3"
)

func emptyHandler(telebot.Context) error {
	return nil
}

func TestContext(t *testing.T) {
	t.Parallel()

	timeout := time.Minute
	baseCtx := context.Background()

	var expectedCtx context.Context

	mockContext := mocks.NewMockContext(t)
	mockContext.
		On("Set", contextKeyContext, mock.Anything).
		Run(func(args mock.Arguments) {
			ctx, ok := args.Get(1).(context.Context)
			require.True(t, ok)

			expectedCtx = ctx
		}).
		Return().
		Once()

	err := WithTimeoutContext(baseCtx, timeout)(emptyHandler)(mockContext)
	require.NoError(t, err)

	mockContext.
		On("Get", contextKeyContext).
		Return(expectedCtx).
		Once()

	ctx := GetContextFromContext(mockContext)
	require.Equal(t, expectedCtx, ctx)
}
