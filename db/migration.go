package db

import (
	"fmt"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DBMigrationUP(dbConfig DBConfig) error {
	// Set up database connection
	db, err := sql.Open("postgres", dbConfig.DBUri)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up migration driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Failed to set up migration driver: %v", err)
	}

	// Set up migration source
	m, err := migrate.NewWithDatabaseInstance(dbConfig.MigrationFilePath, dbConfig.DBName, driver)
	if err != nil {
		return err
	}
	defer m.Close()

	// Apply "up" migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
