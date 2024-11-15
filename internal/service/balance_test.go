package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Weeping-Willow/entain-task/internal/repository"
	mockUserStorage "github.com/Weeping-Willow/entain-task/internal/repository/mocks/UserStorage"
	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
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

func TestPostNewTransaction(t *testing.T) {
	t.Parallel()

	request := spec.PostUserUserIdTransactionRequestObject{
		UserId: 1,
		Body: &spec.Transaction{
			TransactionId: "1",
			Amount:        "100.00",
			State:         spec.Win,
		},
		Params: spec.PostUserUserIdTransactionParams{
			SourceType: spec.Server,
		},
	}

	t.Run("Success positive number", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, request.Body.TransactionId).Return(false, nil)
		mockRepo.EXPECT().UpdateBalanceByAmount(mock.Anything, request.UserId, 100.00, repository.UserTransactionEntity{
			TransactionID: request.Body.TransactionId,
			UserID:        request.UserId,
			Amount:        100.00,
			State:         string(request.Body.State),
			SourceType:    string(request.Params.SourceType),
		}).Return(200.00, nil)

		balance, err := b.PostNewTransaction(context.Background(), request)
		assert.NoError(t, err, "Error should be nil")
		assert.Equal(t, 200.00, balance, "Balance should be 200.00")
	})

	t.Run("Success negative number", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		newRequest := spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Body: &spec.Transaction{
				TransactionId: "1",
				Amount:        "100.00",
				State:         spec.Lose,
			},
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, newRequest.Body.TransactionId).Return(false, nil)
		mockRepo.EXPECT().UpdateBalanceByAmount(mock.Anything, request.UserId, -100.00, repository.UserTransactionEntity{
			TransactionID: newRequest.Body.TransactionId,
			UserID:        newRequest.UserId,
			Amount:        100.00,
			State:         string(newRequest.Body.State),
			SourceType:    string(newRequest.Params.SourceType),
		}).Return(200.00, nil)

		balance, err := b.PostNewTransaction(context.Background(), newRequest)
		assert.NoError(t, err, "Error should be nil")
		assert.Equal(t, 200.00, balance, "Balance should be 200.00")
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(0, repository.ErrUserNotFound)

		balance, err := b.PostNewTransaction(context.Background(), request)
		assert.ErrorIs(t, err, repository.ErrUserNotFound, "Error should be ErrUserNotFound")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})

	t.Run("transactions db error", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, request.Body.TransactionId).Return(false, errors.New("db error"))

		balance, err := b.PostNewTransaction(context.Background(), request)
		assert.ErrorContains(t, err, "db error", "Error should be db error")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})

	t.Run("transactions already present in DB", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, request.Body.TransactionId).Return(true, nil)

		balance, err := b.PostNewTransaction(context.Background(), request)
		assert.ErrorIs(t, err, ErrTransactionAlreadyExists, "Error should be transaction already taken error")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})

	t.Run("transactions already present in DB", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		newRequest := spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Body: &spec.Transaction{
				TransactionId: "1",
				Amount:        "a",
				State:         spec.Lose,
			},
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, newRequest.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, newRequest.Body.TransactionId).Return(false, nil)

		balance, err := b.PostNewTransaction(context.Background(), newRequest)
		assert.ErrorIs(t, err, ErrInvalidAmount, "Error should be invalid amount")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})

	t.Run("not enough balance for the transaction", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		newRequest := spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Body: &spec.Transaction{
				TransactionId: "1",
				Amount:        "100.00",
				State:         spec.Lose,
			},
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, newRequest.UserId).Return(99.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, newRequest.Body.TransactionId).Return(false, nil)

		balance, err := b.PostNewTransaction(context.Background(), newRequest)
		assert.ErrorIs(t, err, ErrNotEnoughBalance, "Error should be not enough balance")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})

	t.Run("update balance by amount fails", func(t *testing.T) {
		mockRepo := mockUserStorage.NewMockUserStorage(t)
		b := balance{
			balanceRepository: mockRepo,
		}

		mockRepo.EXPECT().GetUserBalance(mock.Anything, request.UserId).Return(100.00, nil)
		mockRepo.EXPECT().CheckTransactionExists(mock.Anything, request.Body.TransactionId).Return(false, nil)
		mockRepo.EXPECT().UpdateBalanceByAmount(mock.Anything, request.UserId, 100.00, repository.UserTransactionEntity{
			TransactionID: request.Body.TransactionId,
			UserID:        request.UserId,
			Amount:        100,
			State:         string(request.Body.State),
			SourceType:    string(request.Params.SourceType),
		}).Return(0, errors.New("db error"))

		balance, err := b.PostNewTransaction(context.Background(), request)
		assert.ErrorContains(t, err, "db error", "Error should be db error")
		assert.Equal(t, float64(0), balance, "Balance should be 0")
	})
}
