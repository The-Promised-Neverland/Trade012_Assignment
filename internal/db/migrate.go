package db

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("migration driver error: %w", err)
	}

	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../../")

	migrationsPath := filepath.Join(projectRoot, "internal/db/migrations")

	migrationURL := "file://" + filepath.ToSlash(migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationURL,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration init error: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration run error: %w", err)
	}

	return nil
}
