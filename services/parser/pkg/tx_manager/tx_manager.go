package txmanager

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrFinishTx = errors.New("failed to finish tx")

type TxManager interface {
	Run(
		ctx context.Context,
		options *sql.TxOptions,
		runTx func(ctx context.Context, tx *sql.Tx) error,
	) error
}

type SQLTransactionManager struct {
	db *sql.DB
}

func New(db *sql.DB) *SQLTransactionManager {
	return &SQLTransactionManager{db: db}
}

func (m *SQLTransactionManager) Run(
	ctx context.Context,
	options *sql.TxOptions,
	runTx func(ctx context.Context, tx *sql.Tx) error,
) error {
	var globalErr error

	tx, err := m.db.BeginTx(ctx, options)
	if err != nil {
		return ErrFinishTx
	}
	defer func() {
		if globalErr != nil {
			txErr := tx.Rollback()
			globalErr = errors.Join(ErrFinishTx, txErr)
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			if e, ok := rec.(error); ok {
				globalErr = e
			} else {
				globalErr = errors.New(fmt.Sprintf("%s", rec))
			}
		}
	}()

	if err = runTx(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}
