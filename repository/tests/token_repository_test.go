package token_repo_tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/pratikjethe/go-token-manager/constants"
	"github.com/pratikjethe/go-token-manager/db"
	"github.com/pratikjethe/go-token-manager/models"
	repositories "github.com/pratikjethe/go-token-manager/repository"
	"github.com/pratikjethe/go-token-manager/utils"
)

var testDbConn *sql.DB

func initiatTestDB() (*sql.DB, error) {

	if testDbConn != nil {
		return testDbConn, nil // use same connection
	}

	dBConfig := db.DBConfig{
		DBHost:            "127.0.0.1",
		DBUser:            "postgres",
		DBPassword:        "password",
		DBPort:            "5432",
		DBName:            "token_db",
		DBUri:             "postgresql://postgres:password@127.0.0.1:5432/token_db?sslmode=disable",
		MigrationFilePath: "file://./../../db/migrations",
	}

	dbConn, err := db.InitDB(dBConfig.DBUri) // ONLY USE TESTING DB URL
	if err != nil {
		return dbConn, err

	}

	err = db.DBMigrationUP(dBConfig)
	if err != nil {
		return dbConn, err

	}
	testDbConn = dbConn
	return testDbConn, err

}

func tearTokenTable() error {
	_, err := testDbConn.Exec("TRUNCATE TABLE token_schema.tokens;")
	fmt.Println(err)
	return err

}

func setupTestToken() *models.Token {
	token, _ := utils.GenerateRandomToken(16)
	return &models.Token{
		Token: token,
		// Add other fields if necessary
	}
}

func TestCreateToken(t *testing.T) {

	dbconn, err := initiatTestDB()
	tearTokenTable()
	defer tearTokenTable()

	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Errorf("Expected no error while creating token, got: %v", err)
	}

	if token.ID == 0 {
		t.Errorf("Expected token ID to be set, got 0")
	}
}

func TestGetAvailableToken(t *testing.T) {
	dbconn, err := initiatTestDB()

	tearTokenTable()
	defer tearTokenTable()

	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	// Insert a test token to ensure one is available.
	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	tx, err := dbconn.Begin()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	activeDuration := 60
	expirationDuration := 3600

	resultToken, err := tokenRepo.GetAvailableToken(tx, activeDuration, expirationDuration)
	if err != nil {
		t.Errorf("Expected to retrieve a token, got error: %v", err)
	}

	if resultToken == nil || resultToken.Token != token.Token {
		t.Errorf("Expected to retrieve token %s, got %v", token.Token, resultToken)
	}
}

func TestAssignToken(t *testing.T) {
	dbconn, err := initiatTestDB()
	tearTokenTable()
	defer tearTokenTable()
	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	tx, err := dbconn.Begin()
	if err != nil {
		t.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	err = tokenRepo.AssignToken(tx, token)
	if err != nil {
		t.Errorf("Expected to assign token, got error: %v", err)
	}
	tx.Commit()
}

func TestDeleteToken(t *testing.T) {
	dbconn, err := initiatTestDB()
	tearTokenTable()
	defer tearTokenTable()
	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	err = tokenRepo.DeleteToken(token.Token)
	if err != nil {
		t.Errorf("Expected token to be deleted, got error: %v", err)
	}

	// Try deleting the same token again to test the already deleted scenario.
	err = tokenRepo.DeleteToken(token.Token)
	if err != constants.ERR_TOKEN_ALREADY_DELETED {
		t.Errorf("Expected error %v, got %v", constants.ERR_TOKEN_ALREADY_DELETED, err)
	}

}

func TestUnblockToken(t *testing.T) {
	dbconn, err := initiatTestDB()
	tearTokenTable()
	defer tearTokenTable()
	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	tx, err := dbconn.Begin()
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}
	err = tokenRepo.AssignToken(tx, token)
	if err != nil {
		t.Fatalf("Failed to assign token: %v", err)
		tx.Rollback()
	}
	tx.Commit()
	err = tokenRepo.UnblockToken(token.Token, 3600)
	if err != nil {
		t.Errorf("Expected token to be unblocked, got error: %v", err)
	}
}

func TestKeepAliveToken(t *testing.T) {
	dbconn, err := initiatTestDB()
	tearTokenTable()
	defer tearTokenTable()
	if err != nil {
		t.Errorf("Unexpected DB error %v", err)
	}
	tokenRepo := repositories.NewTokenRepository(dbconn)
	token := setupTestToken()

	err = tokenRepo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	tx, err := dbconn.Begin()
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}
	err = tokenRepo.AssignToken(tx, token)
	if err != nil {
		t.Fatalf("Failed to assign token: %v", err)
		tx.Rollback()
	}
	tx.Commit()

	err = tokenRepo.KeepAliveToken(token.Token, 3600)
	if err != nil {
		t.Errorf("Expected token to be kept alive, got error: %v", err)
	}
}
