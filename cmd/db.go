package main

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	_ "modernc.org/sqlite" // SQLite bindings
)

func newDB(ctx context.Context, file string) (*sql.DB, error) {
	// TODO: migration
	db, err := sql.Open("sqlite", file)
	if err != nil {
		return nil, fmt.Errorf("opening SQLite connection: %w", err)
	}
	if err = initTables(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}

const _sqlCreateSystemsTable = `
	CREATE TABLE IF NOT EXISTS systems (
		symbol TEXT PRIMARY KEY UNIQUE,
		x INTEGER NOT NULL,
		y INTEGER NOT NULL,
		type TEXT NOT NULL,
		factions TEXT NOT NULL
	);

	CREATE INDEX IF NOT EXISTS system_x_pos_index ON systems(x);
	CREATE INDEX IF NOT EXISTS system_y_pos_index ON systems(y);
	CREATE INDEX IF NOT EXISTS system_type_index ON systems(type);
`

func initTables(ctx context.Context, db *sql.DB) error {
	log.Debug("creating SQLite tables")
	_, err := db.ExecContext(ctx, _sqlCreateSystemsTable)
	return err
}
