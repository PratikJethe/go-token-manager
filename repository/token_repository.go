package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pratikjethe/go-token-manager/models"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (repo *TokenRepository) CreateToken(token *models.Token) error {
	query := `
        INSERT INTO token_schema.tokens (token, last_activation_time, is_deleted)
        VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
    `
	err := repo.db.QueryRow(query, token.Token, token.LastActivationTime, token.IsDeleted).
		Scan(&token.ID, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return fmt.Errorf("Error while inserting token: %v", err)
	}
	return nil
}
