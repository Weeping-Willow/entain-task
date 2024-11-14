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
	CheckTransactionExists(ctx context.Context, transactionID string) (bool, error)
	UpdateBalanceByAmount(ctx context.Context, userID uint64, amount float64, entity UserTransactionEntity) (float64, error)
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

func (u *userStorage) UpdateBalanceByAmount(ctx context.Context, userID uint64, amount float64, entity UserTransactionEntity) (float64, error) {
	tx, err := u.db.Beginx()
	if err != nil {
		return 0, errors.Wrap(err, "begin transaction")
	}

	err = u.addToUserBalance(ctx, tx, userID, amount)
	if err != nil {
		slog.Error("update user balance", "error", err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			slog.Error("rollback transaction", "error", txErr.Error())
		}

		return 0, errors.Wrap(err, "update user balance")
	}

	err = u.newTransaction(ctx, tx, entity)
	if err != nil {
		slog.Error("create new transaction", "error", err.Error())
		txErr := tx.Rollback()
		if txErr != nil {
			slog.Error("rollback transaction", "error", txErr.Error())
		}

		return 0, errors.Wrap(err, "create new transaction")
	}

	err = tx.Commit()
	if err != nil {
		return 0, errors.Wrap(err, "update balance transaction")
	}

	return u.GetUserBalance(ctx, userID)
}

func (u *userStorage) newTransaction(ctx context.Context, tx *sqlx.Tx, transaction UserTransactionEntity) error {
	query := `INSERT INTO transactions (id,user_id, amount, state, source_type) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	err := tx.QueryRowxContext(ctx, query, transaction.TransactionID, transaction.UserID, transaction.Amount,
		transaction.State, transaction.SourceType).StructScan(&transaction)
	if err != nil {
		slog.Error("create new transaction", "error", err.Error())

		return errors.Wrap(err, "create new transaction")
	}

	return nil
}

func (u *userStorage) addToUserBalance(ctx context.Context, tx *sqlx.Tx, userID uint64, balance float64) error {
	query := `UPDATE users SET balance = balance + $1 WHERE id = $2`

	_, err := tx.ExecContext(ctx, query, balance, userID)
	if err != nil {
		slog.Error("update user balance", "error", err.Error())

		return errors.Wrap(err, "update user balance")
	}

	return nil
}
