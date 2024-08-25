// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package query

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getSystemByNameStmt, err = db.PrepareContext(ctx, getSystemByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetSystemByName: %w", err)
	}
	if q.getSystemsInRectStmt, err = db.PrepareContext(ctx, getSystemsInRect); err != nil {
		return nil, fmt.Errorf("error preparing query GetSystemsInRect: %w", err)
	}
	if q.hasSystemsRowsStmt, err = db.PrepareContext(ctx, hasSystemsRows); err != nil {
		return nil, fmt.Errorf("error preparing query HasSystemsRows: %w", err)
	}
	if q.insertSystemStmt, err = db.PrepareContext(ctx, insertSystem); err != nil {
		return nil, fmt.Errorf("error preparing query InsertSystem: %w", err)
	}
	if q.insertWaypointStmt, err = db.PrepareContext(ctx, insertWaypoint); err != nil {
		return nil, fmt.Errorf("error preparing query InsertWaypoint: %w", err)
	}
	if q.truncateSystemsStmt, err = db.PrepareContext(ctx, truncateSystems); err != nil {
		return nil, fmt.Errorf("error preparing query TruncateSystems: %w", err)
	}
	if q.truncateWaypointsStmt, err = db.PrepareContext(ctx, truncateWaypoints); err != nil {
		return nil, fmt.Errorf("error preparing query TruncateWaypoints: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getSystemByNameStmt != nil {
		if cerr := q.getSystemByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSystemByNameStmt: %w", cerr)
		}
	}
	if q.getSystemsInRectStmt != nil {
		if cerr := q.getSystemsInRectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSystemsInRectStmt: %w", cerr)
		}
	}
	if q.hasSystemsRowsStmt != nil {
		if cerr := q.hasSystemsRowsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing hasSystemsRowsStmt: %w", cerr)
		}
	}
	if q.insertSystemStmt != nil {
		if cerr := q.insertSystemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertSystemStmt: %w", cerr)
		}
	}
	if q.insertWaypointStmt != nil {
		if cerr := q.insertWaypointStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertWaypointStmt: %w", cerr)
		}
	}
	if q.truncateSystemsStmt != nil {
		if cerr := q.truncateSystemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing truncateSystemsStmt: %w", cerr)
		}
	}
	if q.truncateWaypointsStmt != nil {
		if cerr := q.truncateWaypointsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing truncateWaypointsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	getSystemByNameStmt   *sql.Stmt
	getSystemsInRectStmt  *sql.Stmt
	hasSystemsRowsStmt    *sql.Stmt
	insertSystemStmt      *sql.Stmt
	insertWaypointStmt    *sql.Stmt
	truncateSystemsStmt   *sql.Stmt
	truncateWaypointsStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		getSystemByNameStmt:   q.getSystemByNameStmt,
		getSystemsInRectStmt:  q.getSystemsInRectStmt,
		hasSystemsRowsStmt:    q.hasSystemsRowsStmt,
		insertSystemStmt:      q.insertSystemStmt,
		insertWaypointStmt:    q.insertWaypointStmt,
		truncateSystemsStmt:   q.truncateSystemsStmt,
		truncateWaypointsStmt: q.truncateWaypointsStmt,
	}
}
