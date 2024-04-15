package common

import (
	"bytes"
	"fmt"
	"log/slog"
	"testing"

	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/require"
	"gopkg.in/telebot.v3"
)

func TestLogCommands(t *testing.T) { //nolint:paralleltest  // changing default logger
	output := new(bytes.Buffer)

	logHandler := slog.NewTextHandler(output, nil)
	logger := slog.New(logHandler)

	oldLogger := slog.Default()

	t.Cleanup(func() {
		slog.SetDefault(oldLogger)
	})

	slog.SetDefault(logger)

	expectedArg := "expected-arg"

	mockCtx := mocks.NewMockContext(t)
	mockCtx.On("Args").Return([]string{expectedArg})

	err := LogCommands()(emptyHandler)(mockCtx)
	require.NoError(t, err)

	require.Contains(t, output.String(), expectedArg)
}

func TestLogCommandsErr(t *testing.T) { //nolint:paralleltest  // changing default logger
	output := new(bytes.Buffer)

	logHandler := slog.NewTextHandler(output, nil)
	logger := slog.New(logHandler)

	oldLogger := slog.Default()

	t.Cleanup(func() {
		slog.SetDefault(oldLogger)
	})

	slog.SetDefault(logger)

	expectedArg := "expected-arg2"

	mockCtx := mocks.NewMockContext(t)
	mockCtx.On("Args").Return([]string{expectedArg})

	expectedErr := fmt.Errorf("handler error")

	err := LogCommands()(func(telebot.Context) error {
		return expectedErr
	})(mockCtx)
	require.ErrorIs(t, err, expectedErr)

	require.Contains(t, output.String(), expectedArg)
	require.Contains(t, output.String(), "error")
}
