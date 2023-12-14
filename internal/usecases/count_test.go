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

func TestCount(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	player := &entities.Player{
		ID:         uuid.New(),
		TelegramID: rand.Int63(),
	}

	t.Run("UpdateCountByTelegramID", func(t *testing.T) {
		t.Parallel()

		mockRepo := mocks.NewMockrepository(t)

		cases := new(UseCases)

		cases.WithRepo(mockRepo)

		countDelta := int16(rand.Intn(0xff))

		mockRepo.On("FindUserByTelegramID", ctx, player.TelegramID).Return(player, nil).Once()
		mockRepo.On("InsertPlayerCountChange", ctx, player.ID, countDelta).Return(nil).Once()

		err := cases.UpdateCountByTelegramID(ctx, player.TelegramID, countDelta)
		require.NoError(t, err)
	})

	t.Run("GetPlayerCountByTelegramID", func(t *testing.T) {
		t.Parallel()

		mockRepo := mocks.NewMockrepository(t)
		cases := new(UseCases)

		cases.WithRepo(mockRepo)

		expectedCount := rand.Int31()

		mockRepo.
			On("FindUserByTelegramID", ctx, player.TelegramID).
			Return(player, nil).Once()
		mockRepo.
			On("GetPlayerCount", ctx, player.ID).
			Return(expectedCount, nil).Once()

		count, err := cases.GetPlayerCountByTelegramID(ctx, player.TelegramID)
		require.NoError(t, err)
		require.Equal(t, expectedCount, count)
	})
}
