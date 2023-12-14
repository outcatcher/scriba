package repo

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
	"github.com/outcatcher/scriba/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRepo_GetPlayerCount_emptyCache(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()
	expectedCount := rand.Int31()

	mockExecutor.
		On("GetContext",
			ctx,
			mock.AnythingOfType("*int32"),
			mock.AnythingOfType("string"),
			playerUUID).
		Run(func(args mock.Arguments) {
			count, ok := args.Get(1).(*int32)
			require.True(t, ok)

			*count = expectedCount
		}).
		Return(nil)

	count, err := repo.GetPlayerCount(ctx, playerUUID)
	require.NoError(t, err)
	require.Equal(t, expectedCount, count)

	cachedCount, ok := repo.countSumCache.Get(playerUUID)
	require.True(t, ok)
	require.Equal(t, expectedCount, cachedCount)
}

func TestRepo_GetPlayerCount_cached(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerID := uuid.New()
	expectedCount := rand.Int31()

	repo.countSumCache.Set(playerID, expectedCount)

	count, err := repo.GetPlayerCount(ctx, playerID)
	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}

func TestRepo_GetPlayerCount_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()

	mockExecutor.
		On("GetContext",
			ctx,
			mock.AnythingOfType("*int32"),
			mock.AnythingOfType("string"),
			playerUUID).
		Return(errTest)

	_, err := repo.GetPlayerCount(ctx, playerUUID)
	require.ErrorIs(t, err, errTest)
}

func TestRepo_InsertPlayerCountChange_ok(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()
	expectedDelta := int16(rand.Int31n(math.MaxInt16))

	repo.countSumCache.Set(playerUUID, 0xff)

	mockExecutor.On("ExecContext", ctx, mock.AnythingOfType("string"), expectedDelta, playerUUID).Return(nil, nil)

	err := repo.InsertPlayerCountChange(ctx, playerUUID, expectedDelta)
	require.NoError(t, err)

	_, ok := repo.countSumCache.Get(playerUUID) // cache should be reset for the user
	require.False(t, ok)
}

func TestRepo_InsertPlayerCountChange_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()
	expectedDelta := int16(rand.Int31n(math.MaxInt16))

	repo.countSumCache.Set(playerUUID, 0xff)

	mockExecutor.
		On("ExecContext", ctx, mock.AnythingOfType("string"), expectedDelta, playerUUID).
		Return(nil, errTest)

	err := repo.InsertPlayerCountChange(ctx, playerUUID, expectedDelta)
	require.ErrorIs(t, err, errTest)

	_, ok := repo.countSumCache.Get(playerUUID) // cache should NOT be reset for the user
	require.True(t, ok)
}

func TestGetCountHistoryForPeriod_ok(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()
	endDate := time.Now()
	startDate := endDate.Add(-24 * time.Hour)

	expectedHistory := []entities.CountHistoryEvent{{
		Timestamp: endDate.Add(-2 * time.Hour),
		Delta:     12,
	}}

	mockedHistory := []countHistoryItem[entities.CountHistoryEvent]{{
		Timestamp: endDate.Add(-2 * time.Hour),
		Delta:     12,
	}}

	mockExecutor.On(
		"SelectContext",
		ctx,
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("uuid.UUID"),
		startDate,
		endDate,
	).Run(func(args mock.Arguments) {
		items, ok := args.Get(1).(*[]countHistoryItem[entities.CountHistoryEvent])
		require.True(t, ok)

		*items = mockedHistory
	}).Return(nil)

	history, err := repo.GetCountHistoryForPeriod(ctx, playerUUID, startDate, endDate)
	require.NoError(t, err)
	require.Equal(t, expectedHistory, history)
}

func TestGetCountHistoryForPeriod_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	playerUUID := uuid.New()
	endDate := time.Now()
	startDate := endDate.Add(-24 * time.Hour)

	mockExecutor.On(
		"SelectContext",
		ctx,
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("uuid.UUID"),
		startDate,
		endDate,
	).Return(errTest)

	_, err := repo.GetCountHistoryForPeriod(ctx, playerUUID, startDate, endDate)
	require.ErrorIs(t, err, errTest)
}
