package service

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Weeping-Willow/entain-task/internal/repository"
	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

type Balance interface {
	GetUserBalance(ctx context.Context, userID uint64) (float64, error)
	PostNewTransaction(ctx context.Context, request spec.PostUserUserIdTransactionRequestObject) (float64, error)
}

type balance struct {
	balanceRepository repository.UserStorage
}

var (
	ErrParseTransactionAmount   = errors.New("parse transaction amount")
	ErrTransactionAlreadyExists = errors.New("transaction already exists")
	ErrNotEnoughBalance         = errors.New("not enough balance")
)

func NewBalance(balanceRepository repository.UserStorage) Balance {
	return &balance{
		balanceRepository: balanceRepository,
	}
}

func (b *balance) GetUserBalance(ctx context.Context, userID uint64) (float64, error) {
	balanceAvail, err := b.balanceRepository.GetUserBalance(ctx, userID)

	return balanceAvail, errors.Wrap(err, "get user balance")
}

func (b *balance) PostNewTransaction(ctx context.Context, request spec.PostUserUserIdTransactionRequestObject) (float64, error) {
	// Lock Transaction
	// Lock User Balance update

	userBalance, err := b.GetUserBalance(ctx, request.UserId)
	if err != nil {
		return 0, err
	}

	transactionExists, err := b.balanceRepository.CheckTransactionExists(ctx, request.Body.TransactionId)
	if err != nil {
		return 0, errors.Wrap(err, "check transaction exists")
	}

	if transactionExists {
		return 0, ErrTransactionAlreadyExists
	}

	transactionAmount, err := strconv.ParseFloat(request.Body.Amount, 64)
	if err != nil {
		return 0, ErrParseTransactionAmount
	}

	if request.Body.State == spec.Lose {
		if userBalance < transactionAmount {
			return 0, ErrNotEnoughBalance
		}

		transactionAmount = transactionAmount * -1
	}

	newBalance, err := b.balanceRepository.UpdateBalanceByAmount(ctx, request.UserId, transactionAmount, repository.UserTransactionEntity{
		TransactionID: request.Body.TransactionId,
		UserID:        request.UserId,
		Amount:        transactionAmount,
		State:         string(request.Body.State),
		SourceType:    string(request.Params.SourceType),
	})
	if err != nil {
		return 0, errors.Wrap(err, "update user balance")
	}

	// remove user lock

	return newBalance, nil
}
