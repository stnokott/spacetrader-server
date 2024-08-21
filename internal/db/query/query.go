// Package query contains the sqlc-generated functions and some manually-created utility functions.
package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Tx wraps the sqlc-generated queries into a transaction.
type Tx struct {
	*Queries

	tx *sql.Tx
}

// WithTx wraps the sqlc-generated queries into a transaction.
//
// The caller should call Done() when the transaction is no longer needed.
func WithTx(ctx context.Context, db *sql.DB, q *Queries) (Tx, error) {
	log.Debug("creating transaction")
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return Tx{}, fmt.Errorf("creating transaction: %w", err)
	}

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
		log.Debug("rolling transaction back")
		errRollback := t.tx.Rollback()
		if errRollback != nil {
			log.Errorf("failed to rollback: %v", errRollback)
		}
		return errors.Join(err, errRollback)
	}

	log.Debug("committing transaction")
	if errCommit := t.tx.Commit(); errCommit != nil {
		log.Errorf("failed to commit: %v", errCommit)
		return errCommit
	}
	return nil
}
