package usecases

import (
	"context"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/outcatcher/scriba/internal/entities"
	"github.com/outcatcher/scriba/internal/usecases/mocks"
	"github.com/stretchr/testify/require"
)

func TestRegisterWithTelegram(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	mockRepo := mocks.NewMockrepository(t)
	cases := new(UseCases)

	cases.WithRepo(mockRepo)

	expectedID := rand.Int63()

	mockRepo.On("CreateUserFromTG", ctx, expectedID).Return(uuid.New(), nil).Once()

	err := cases.RegisterWithTelegram(ctx, expectedID)
	require.NoError(t, err)
}

func TestListPlayers(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	mockRepo := mocks.NewMockrepository(t)
	cases := new(UseCases)

	cases.WithRepo(mockRepo)

	players := []entities.Player{{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}}

	mockRepo.On("ListPlayers", ctx).Return(players, nil).Once()

	entPlayers, err := cases.ListPlayers(ctx)
	require.NoError(t, err)

	require.Len(t, entPlayers, 1)
	require.Equal(t, players[0].TelegramID, entPlayers[0].TelegramID)
	require.Equal(t, players[0].ID, entPlayers[0].ID)
}
