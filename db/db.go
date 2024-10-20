package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBUri             string
	MigrationFilePath string
}

func GetConfig() (*DBConfig, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbMigrationPath := os.Getenv("MIGRATION_FILE_PATH")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbMigrationPath == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return &DBConfig{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBUri:      dbURI,
		MigrationFilePath: fmt.Sprintf("file://%s",
			dbMigrationPath),
	}, nil
}

func InitDB(dbURI string) (*sql.DB, error) {
	var err error

	var db *sql.DB
	db, err = sql.Open("postgres", dbURI)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err

	}

	return db, err
}
