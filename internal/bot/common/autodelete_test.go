package common

import (
	"testing"
	"time"

	"github.com/outcatcher/scriba/internal/bot/schema/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAutodeleteMiddleware(t *testing.T) {
	t.Parallel()

	expectedTime := time.Millisecond
	mockCtx := mocks.NewMockContext(t)

	mockCtx.
		On("DeleteAfter", expectedTime).
		Return(nil) // not using returned timer

	err := AutodeleteMiddleware(expectedTime)(emptyHandler)(mockCtx)
	require.NoError(t, err)
}

func TestAutodeleteMiddlewareNoTimeout(t *testing.T) {
	t.Parallel()

	mockCtx := new(mocks.MockContext) // methods won't be called, can't use NewMockContext

	mockCtx.
		On("DeleteAfter", mock.Anything).
		Return(nil) // not using returned timer

	err := AutodeleteMiddleware(0)(emptyHandler)(mockCtx)
	require.NoError(t, err)

	mockCtx.AssertNotCalled(t, "DeleteAfter")
}
