package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Weeping-Willow/entain-task/internal/repository"
	mockUserStorage "github.com/Weeping-Willow/entain-task/internal/repository/mocks/UserStorage"
)

func TestGetUserBalance(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		userID := uint64(1)

		mockRepo.EXPECT().GetUserBalance(mock.Anything, userID).Return(100.00, nil)

		balanceService := NewBalance(mockRepo)
		balance, err := balanceService.GetUserBalance(context.Background(), userID)

		assert.NoError(t, err, "Error should be nil")
		assert.Equal(t, 100.00, balance, "Balance should be 100.00")
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		userID := uint64(1)

		mockRepo.EXPECT().GetUserBalance(mock.Anything, userID).Return(0, repository.ErrUserNotFound)

		balanceService := NewBalance(mockRepo)
		balance, err := balanceService.GetUserBalance(context.Background(), userID)

		assert.ErrorIs(t, err, repository.ErrUserNotFound, "Error should be ErrUserNotFound")
		assert.Empty(t, 0, balance, "Balance should be 0")
	})
}
