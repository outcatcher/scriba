package repo

import (
	"context"
	"errors"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
	"github.com/outcatcher/scriba/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var errTest = errors.New("db test error")

func TestRepo_CreateUserFromTG_ok(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedID := int64(1)
	expectedUUID := uuid.New()

	repo.usersCache = make([]entities.Player, 10)

	mockExecutor.
		On("GetContext", ctx, mock.AnythingOfType("*uuid.UUID"), mock.AnythingOfType("string"), expectedID).
		Run(func(args mock.Arguments) {
			createdID, ok := args.Get(1).(*uuid.UUID)
			require.True(t, ok)

			*createdID = expectedUUID
		}).Return(nil).Once()

	actualUUID, err := repo.CreateUserFromTG(ctx, expectedID)
	require.NoError(t, err)

	require.Equal(t, expectedUUID, actualUUID)
	require.Empty(t, repo.usersCache)
}

func TestRepo_CreateUserFromTG_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedID := int64(1)

	repo.usersCache = make([]entities.Player, 10)

	mockExecutor.
		On("GetContext", ctx, mock.AnythingOfType("*uuid.UUID"), mock.AnythingOfType("string"), expectedID).
		Return(errTest).Once()

	_, err := repo.CreateUserFromTG(ctx, expectedID)
	require.ErrorIs(t, err, errTest)

	require.NotEmpty(t, repo.usersCache) // check cache is not emptied
}

func TestRepo_ListPlayers_emptyCache(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedPlayers := []entities.Player{{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}}

	mockExecutor.
		On("SelectContext", ctx, mock.Anything, mock.AnythingOfType("string")).
		Run(func(args mock.Arguments) {
			playerList, ok := args.Get(1).(*[]player)
			require.True(t, ok)

			*playerList = append(*playerList, player{
				ID:       expectedPlayers[0].ID,
				TGUserID: expectedPlayers[0].TelegramID,
			})
		}).Return(nil)

	players, err := repo.ListPlayers(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedPlayers, players)
	require.Equal(t, expectedPlayers, repo.usersCache)
}

func TestRepo_ListPlayers_cached(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedPlayers := []entities.Player{{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}}

	repo.usersCache = expectedPlayers

	players, err := repo.ListPlayers(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedPlayers, players)
}

func TestRepo_ListPlayers_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	mockExecutor.
		On("SelectContext", ctx, mock.Anything, mock.AnythingOfType("string")).Return(errTest)

	_, err := repo.ListPlayers(ctx)
	require.ErrorIs(t, err, errTest)
}

func TestRepo_FindUserByTelegramID_emptyCache(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedPlayer := &entities.Player{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}

	mockExecutor.
		On("GetContext",
			ctx, mock.AnythingOfType("*repo.player"), mock.AnythingOfType("string"), expectedPlayer.TelegramID).
		Run(func(args mock.Arguments) {
			pl, ok := args.Get(1).(*player)
			require.True(t, ok)

			pl.TGUserID = expectedPlayer.TelegramID
			pl.ID = expectedPlayer.ID
		}).
		Return(nil).Once()

	playerUser, err := repo.FindUserByTelegramID(ctx, expectedPlayer.TelegramID)
	require.NoError(t, err)
	require.Equal(t, expectedPlayer, playerUser)

	cached, ok := repo.usersCacheByTGID.Get(expectedPlayer.TelegramID)
	require.True(t, ok)
	require.Equal(t, expectedPlayer, cached)
}

func TestRepo_FindUserByTelegramID_cached(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()

	expectedPlayer := &entities.Player{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}

	repo.usersCacheByTGID.Set(expectedPlayer.TelegramID, expectedPlayer)

	playerUser, err := repo.FindUserByTelegramID(ctx, expectedPlayer.TelegramID)
	require.NoError(t, err)
	require.Equal(t, expectedPlayer, playerUser)
}

func TestRepo_FindUserByTelegramID_dbError(t *testing.T) {
	t.Parallel()

	mockExecutor := mocks.NewMockqueryExecutor(t)

	repo := New(mockExecutor)
	ctx := context.Background()
	randomID := rand.Int63()

	mockExecutor.
		On("GetContext",
			ctx, mock.AnythingOfType("*repo.player"), mock.AnythingOfType("string"), randomID).
		Return(errTest).Once()

	expectedPlayer := &entities.Player{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}

	repo.usersCacheByTGID.Set(expectedPlayer.TelegramID, expectedPlayer)

	_, err := repo.FindUserByTelegramID(ctx, randomID)
	require.ErrorIs(t, err, errTest)

	cached, ok := repo.usersCacheByTGID.Get(expectedPlayer.TelegramID)
	require.True(t, ok)
	require.Equal(t, expectedPlayer, cached)
}
