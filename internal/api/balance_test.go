package api

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Weeping-Willow/entain-task/internal/config"
	"github.com/Weeping-Willow/entain-task/internal/repository"
	"github.com/Weeping-Willow/entain-task/internal/service"
	mockBalance "github.com/Weeping-Willow/entain-task/internal/service/mocks/Balance"
	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

func TestGetUserUserIdBalance(t *testing.T) {
	t.Parallel()

	t.Run("service returns user not found", func(t *testing.T) {
		userID := uint64(2)
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := Server{
			balanceService: mockBalanceService,
		}

		mockBalanceService.EXPECT().GetUserBalance(mock.Anything, userID).Return(0, errors.Wrap(repository.ErrUserNotFound, "get user balance"))

		res, err := s.GetUserUserIdBalance(context.Background(), spec.GetUserUserIdBalanceRequestObject{UserId: userID})
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorIs(t, err, repository.ErrUserNotFound, "Error should be ErrUserNotFound")
	})

	t.Run("happy", func(t *testing.T) {
		userID := uint64(2)
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := Server{
			balanceService: mockBalanceService,
		}

		mockBalanceService.EXPECT().GetUserBalance(mock.Anything, userID).Return(21, nil)

		res, err := s.GetUserUserIdBalance(context.Background(), spec.GetUserUserIdBalanceRequestObject{UserId: userID})
		assert.Equal(t, res, spec.GetUserUserIdBalance200JSONResponse{
			UserId:  userID,
			Balance: "21.00",
		})
		assert.Empty(t, err, nil, "Error should be empty")
	})
}

func TestPostUserUserIdTransaction(t *testing.T) {
	t.Parallel()

	t.Run("validation error: source type", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)

		res, err := s.PostUserUserIdTransaction(context.Background(), spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: "invalid",
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount:        "22",
				State:         spec.Win,
				TransactionId: "1",
			},
		})
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorContains(t, err, "Field validation for 'SourceType' failed on the 'oneof' tag", "Error should be validation error")
	})

	t.Run("validation error: state", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)

		res, err := s.PostUserUserIdTransaction(context.Background(), spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount:        "22",
				State:         "invalid",
				TransactionId: "1",
			},
		})
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorContains(t, err, "Field validation for 'State' failed on the 'oneof' tag", "Error should be validation error")
	})

	t.Run("validation error: TransactionId", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)

		res, err := s.PostUserUserIdTransaction(context.Background(), spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount: "22",
				State:  spec.Lose,
			},
		})
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorContains(t, err, "Field validation for 'TransactionId' failed on the 'required' tag")
	})

	t.Run("validation error: Amount", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)

		res, err := s.PostUserUserIdTransaction(context.Background(), spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount:        "a",
				State:         spec.Lose,
				TransactionId: "1",
			},
		})
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorContains(t, err, "Field validation for 'Amount' failed on the 'numeric' tag")
	})

	t.Run("service returns error", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)
		req := spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount:        "22",
				State:         spec.Lose,
				TransactionId: "1",
			},
		}

		mockBalanceService.EXPECT().PostNewTransaction(mock.Anything, req).Return(0, service.ErrNotEnoughBalance)

		res, err := s.PostUserUserIdTransaction(context.Background(), req)
		assert.Empty(t, res, "Response should be empty")
		assert.ErrorIs(t, err, service.ErrNotEnoughBalance, "Error should be balance error")
	})

	t.Run("happy", func(t *testing.T) {
		mockBalanceService := mockBalance.NewMockBalance(t)
		s := New(config.Config{}, mockBalanceService)
		req := spec.PostUserUserIdTransactionRequestObject{
			UserId: 1,
			Params: spec.PostUserUserIdTransactionParams{
				SourceType: spec.Server,
			},
			Body: &spec.PostUserUserIdTransactionJSONRequestBody{
				Amount:        "22",
				State:         spec.Lose,
				TransactionId: "1",
			},
		}

		mockBalanceService.EXPECT().PostNewTransaction(mock.Anything, req).Return(24.551, nil)

		res, err := s.PostUserUserIdTransaction(context.Background(), req)
		assert.Equal(t, res, spec.PostUserUserIdTransaction200JSONResponse{
			UserId:  1,
			Balance: "24.55",
		})
		assert.Empty(t, err, nil, "Error should be empty")
	})
}
