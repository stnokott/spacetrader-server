package main

import (
	"database/sql"
	"fmt"

	"github.com/stnokott/spacetrader-server/internal/db"
	_ "modernc.org/sqlite" // SQLite bindings
)

func newDB(file string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", file)
	if err != nil {
		return nil, fmt.Errorf("opening SQLite connection: %w", err)
	}
	if err = db.MigrateUp(conn); err != nil {
		return nil, err
	}
	return conn, nil
}
