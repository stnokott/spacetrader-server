package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite" // SQLite bindings
	"github.com/golang-migrate/migrate/v4/source/iofs"     // Read migrations from embedded FS
)

//go:embed schema/*.sql
var migrationsFS embed.FS

// MigrateUp applies the required migration SQLs to get to the latest schema.
func MigrateUp(conn *sql.DB) error {
	d, err := iofs.New(migrationsFS, "schema")
	if err != nil {
		return fmt.Errorf("opening migrations directory: %w", err)
	}

	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("creating SQLite driver from connection: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", d, "spacetrader", driver)
	if err != nil {
		return fmt.Errorf("creating migration instance: %w", err)
	}
	return m.Up()
}
