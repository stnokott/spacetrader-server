// Package query contains the sqlc-generated functions and some manually-created utility functions.
package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/stnokott/spacetrader-server/internal/log"
)

var logger = log.ForComponent("db")

// Tx wraps the sqlc-generated queries into a transaction.
type Tx struct {
	*Queries

	tx *sql.Tx
}

// WithTx wraps the sqlc-generated queries into a transaction.
//
// The caller should call Done() when the transaction is no longer needed.
func WithTx(ctx context.Context, db *sql.DB, q *Queries) (Tx, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return Tx{}, fmt.Errorf("creating transaction: %w", err)
	}
	logger.Debug("transaction created")

	return Tx{
		Queries: q.WithTx(tx),
		tx:      tx,
	}, nil
}

// Done should be called after the caller has finished operation on the transaction.
//
// Depending on err, it will either perform a commit (when err is nil) or a rollback otherwise.
// If the commit fails, an error will be returned.
// If the rollback fails, the original error, joined with the rollback error, will be returned.
func (t Tx) Done(err error) error {
	if err != nil {
		logger.Debug("rolling transaction back")
		errRollback := t.tx.Rollback()
		if errRollback != nil && !errors.Is(errRollback, sql.ErrTxDone) {
			logger.Errorf("failed to rollback: %v", errRollback)
		}
		return errors.Join(err, errRollback)
	}

	logger.Debug("committing transaction")
	if errCommit := t.tx.Commit(); errCommit != nil {
		logger.Errorf("failed to commit: %v", errCommit)
		return errCommit
	}
	return nil
}
