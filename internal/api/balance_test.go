package api

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Weeping-Willow/entain-task/internal/repository"
	mockBalance "github.com/Weeping-Willow/entain-task/internal/service/mocks/Balance"
	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

func TestGetUserUserIdBalance(t *testing.T) {
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
