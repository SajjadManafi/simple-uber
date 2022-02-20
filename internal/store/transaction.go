package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SajjadManafi/simple-uber/models"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type TXStore struct {
	db DBTX
}

func NewTXStore(tx *sql.Tx) *TXStore {
	return &TXStore{
		db: tx,
	}
}

func (q *PostgresStore) execTx(ctx context.Context, fn func(*TXStore) error) error {

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	t := NewTXStore(tx)

	err = fn(t)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (q *PostgresStore) TripDoneTx(ctx context.Context, arg models.TripDoneTransactionParams) (models.TripDoneTransactionResult, error) {
	var result models.TripDoneTransactionResult

	err := q.execTx(ctx, func(t *TXStore) error {
		var err error

		TUarg := models.UpdateTripDoneParams{
			ID:           arg.ID,
			DriverRating: arg.DriverRating,
		}
		result.Trip, err = q.UpdateTripDone(ctx, TUarg)
		if err != nil {
			return err
		}

		result.User, result.Driver, err = addMoney(ctx, q, result.Trip.RiderID, -result.Trip.Amount, result.Trip.DriverID, result.Trip.Amount)

		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *PostgresStore,
	RiderID int32,
	amount1 int64,
	DriverID int32,
	amount2 int64,
) (Rider models.User, Driver models.Driver, err error) {
	Rider, err = q.AddUserBalance(ctx, models.AddUserBalanceParams{
		ID:     RiderID,
		Amount: amount1,
	})

	if err != nil {
		return
	}

	Driver, err = q.AddDriverBalance(ctx, models.AddDriverBalanceParams{
		ID:     DriverID,
		Amount: amount2,
	})

	return
}
