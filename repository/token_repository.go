package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pratikjethe/go-token-manager/constants"
	"github.com/pratikjethe/go-token-manager/models"
)

type TokenRepository struct {
	DB *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{DB: db}
}

func (repo *TokenRepository) CreateToken(token *models.Token) error {
	query := `
        INSERT INTO token_schema.tokens (token)
        VALUES ($1) RETURNING id, created_at, updated_at
    `
	err := repo.DB.QueryRow(query, token.Token).
		Scan(&token.ID, &token.CreatedAt, &token.UpdatedAt)
	if err != nil {
		return fmt.Errorf("Error while inserting token: %v", err)
	}
	return nil
}

func (r *TokenRepository) GetAvailableToken(tx *sql.Tx, activeTokenDuration int, tokenExpireDuration int) (*models.Token, error) {
	var token models.Token
	query := `
	SELECT id, token, last_activation_time, is_deleted, created_at, updated_at
	FROM token_schema.tokens
	WHERE is_deleted = false
	AND (
		last_activation_time IS NULL 
		OR NOW() BETWEEN last_activation_time +  make_interval(secs => $1) AND last_activation_time +  make_interval(secs => $2)
	)
	LIMIT 1
	FOR UPDATE SKIP LOCKED;`

	err := tx.QueryRow(query, activeTokenDuration, tokenExpireDuration).Scan(
		&token.ID,
		&token.Token,
		&token.LastActivationTime,
		&token.IsDeleted,
		&token.CreatedAt,
		&token.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, constants.ERR_NO_TOKENS
		}
		return nil, err
	}
	return &token, nil
}

func (r *TokenRepository) AssignToken(tx *sql.Tx, token *models.Token) error {
	query := `UPDATE token_schema.tokens SET last_activation_time = $1 WHERE id = $2`
	_, err := tx.Exec(query, time.Now(), token.ID)
	return err
}
