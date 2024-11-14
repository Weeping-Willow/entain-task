package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserBalanceEntity struct {
	Balance float64 `db:"balance"`
}

type UserStorage interface {
	GetUserBalance(ctx context.Context, userID uint64) (float64, error)
}

type balance struct {
	db *sqlx.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserStorage(db *sqlx.DB) UserStorage {
	return &balance{
		db: db,
	}
}

func (b *balance) GetUserBalance(ctx context.Context, userID uint64) (float64, error) {
	query := `SELECT balance FROM users WHERE id = $1`

	var userBalance UserBalanceEntity
	err := b.db.GetContext(ctx, &userBalance, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrUserNotFound
		}

		slog.Error("get user balance from storage", "error", err.Error())

		return 0, errors.Wrap(err, "get user balance from storage")
	}

	return userBalance.Balance, nil
}
