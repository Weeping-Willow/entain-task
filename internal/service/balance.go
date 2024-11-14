package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Weeping-Willow/entain-task/internal/repository"
)

type Balance interface {
	GetUserBalance(ctx context.Context, userID uint64) (float64, error)
}

type balance struct {
	balanceRepository repository.UserStorage
}

func NewBalance(balanceRepository repository.UserStorage) Balance {
	return &balance{
		balanceRepository: balanceRepository,
	}
}

func (b *balance) GetUserBalance(ctx context.Context, userID uint64) (float64, error) {
	balanceAvail, err := b.balanceRepository.GetUserBalance(ctx, userID)

	return balanceAvail, errors.Wrap(err, "get user balance")
}
