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

type UserTransactionEntity struct {
	TransactionID string  `db:"id"`
	UserID        uint64  `db:"user_id"`
	Amount        float64 `db:"amount"`
	State         string  `db:"state"`
	SourceType    string  `db:"source_type"`
}

type UserStorage interface {
	GetUserBalance(ctx context.Context, userID uint64) (float64, error)
	UpdateUserBalance(ctx context.Context, userID uint64, balance float64) error
	NewTransaction(ctx context.Context, transaction UserTransactionEntity) error
	CheckTransactionExists(ctx context.Context, transactionID string) (bool, error)
}

type userStorage struct {
	db *sqlx.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserStorage(db *sqlx.DB) UserStorage {
	return &userStorage{
		db: db,
	}
}

func (u *userStorage) GetUserBalance(ctx context.Context, userID uint64) (float64, error) {
	query := `SELECT balance FROM users WHERE id = $1`

	var userBalance UserBalanceEntity
	err := u.db.GetContext(ctx, &userBalance, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrUserNotFound
		}

		slog.Error("get user balance from storage", "error", err.Error())

		return 0, errors.Wrap(err, "get user balance from storage")
	}

	return userBalance.Balance, nil
}

func (u *userStorage) NewTransaction(ctx context.Context, transaction UserTransactionEntity) error {
	query := `INSERT INTO transactions (id,user_id, amount, state, source_type) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	err := u.db.QueryRowxContext(ctx, query, transaction.TransactionID, transaction.UserID, transaction.Amount,
		transaction.State, transaction.SourceType).StructScan(&transaction)
	if err != nil {
		slog.Error("create new transaction", "error", err.Error())

		return errors.Wrap(err, "create new transaction")
	}

	return nil
}

func (u *userStorage) UpdateUserBalance(ctx context.Context, userID uint64, balance float64) error {
	query := `UPDATE users SET balance = $1 WHERE id = $2`

	_, err := u.db.ExecContext(ctx, query, balance, userID)
	if err != nil {
		slog.Error("update user balance", "error", err.Error())

		return errors.Wrap(err, "update user balance")
	}

	return nil
}

func (u *userStorage) CheckTransactionExists(ctx context.Context, transactionID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM transactions WHERE id = $1)`

	var exists bool
	err := u.db.GetContext(ctx, &exists, query, transactionID)
	if err != nil {
		slog.Error("check transaction exists", "error", err.Error())

		return false, errors.Wrap(err, "check transaction exists")
	}

	return exists, nil
}
