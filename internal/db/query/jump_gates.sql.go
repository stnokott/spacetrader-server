// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: jump_gates.sql

package query

import (
	"context"
)

const getJumpgatesInSystem = `-- name: GetJumpgatesInSystem :many
SELECT
	waypoint, connects_to
FROM jump_gates
WHERE waypoint IN (
	SELECT symbol
	FROM waypoints
	WHERE system = ?1
)
`

func (q *Queries) GetJumpgatesInSystem(ctx context.Context, systemID string) ([]JumpGate, error) {
	rows, err := q.query(ctx, q.getJumpgatesInSystemStmt, getJumpgatesInSystem, systemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []JumpGate{}
	for rows.Next() {
		var i JumpGate
		if err := rows.Scan(&i.Waypoint, &i.ConnectsTo); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const hasJumpgateRows = `-- name: HasJumpgateRows :one
SELECT EXISTS (SELECT 1 FROM jump_gates) AS "exists"
`

func (q *Queries) HasJumpgateRows(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.hasJumpgateRowsStmt, hasJumpgateRows)
	var exists int64
	err := row.Scan(&exists)
	return exists, err
}

const insertJumpGate = `-- name: InsertJumpGate :exec
INSERT INTO jump_gates (
	waypoint, connects_to
) VALUES (
	?, ?
)
`

type InsertJumpGateParams struct {
	Waypoint   string
	ConnectsTo string
}

func (q *Queries) InsertJumpGate(ctx context.Context, arg InsertJumpGateParams) error {
	_, err := q.exec(ctx, q.insertJumpGateStmt, insertJumpGate, arg.Waypoint, arg.ConnectsTo)
	return err
}

const truncateJumpGates = `-- name: TruncateJumpGates :exec
DELETE FROM jump_gates
`

func (q *Queries) TruncateJumpGates(ctx context.Context) error {
	_, err := q.exec(ctx, q.truncateJumpGatesStmt, truncateJumpGates)
	return err
}
