package repositories

import (
	"database/sql"
	"log"
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
		log.Println(err.Error())
		return constants.DB_OPERATION_ERR
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
		log.Println(err.Error())
		if err == sql.ErrNoRows {
			return nil, constants.ERR_NO_TOKENS
		}
		return nil, constants.DB_OPERATION_ERR
	}
	return &token, nil
}

func (r *TokenRepository) AssignToken(tx *sql.Tx, token *models.Token) error {
	query := `UPDATE token_schema.tokens SET last_activation_time = $1 WHERE id = $2`
	_, err := tx.Exec(query, time.Now(), token.ID)
	return err
}

func (r *TokenRepository) DeleteToken(token string) error {
	query := `
		UPDATE token_schema.tokens 
		SET is_deleted = true 
		WHERE token = $1 AND is_deleted = false
	`

	res, err := r.DB.Exec(query, token)
	if err != nil {
		log.Println(err.Error())
		return constants.DB_OPERATION_ERR
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return constants.DB_OPERATION_ERR
	}

	if rowsAffected == 0 {
		return constants.ERR_TOKEN_ALREADY_DELETED
	}

	return nil
}

func (r *TokenRepository) UnblockToken(token string, activeDuration int) error {
	// Create a transaction to ensure atomicity
	query := `
		UPDATE token_schema.tokens
		SET last_activation_time = NULL
		WHERE token = $1 
		AND is_deleted != true 
		AND last_activation_time IS NOT NULL
		AND last_activation_time + make_interval(secs => $2) > NOW()
		RETURNING id;
	`

	var id int
	err := r.DB.QueryRow(query, token, activeDuration).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		if err == sql.ErrNoRows {
			return constants.TOKEN_UNBLOCK_ERR
		}
		return constants.DB_OPERATION_ERR
	}

	return nil
}

func (r *TokenRepository) KeepAliveToken(token string, expirationDuration int) error {
	// Create a transaction to ensure atomicity
	query := `
		UPDATE token_schema.tokens
		SET last_activation_time = NOW()
		WHERE token = $1 
		AND is_deleted != true 
		AND last_activation_time IS NOT NULL
		AND last_activation_time + make_interval(secs => $2) > NOW()
		RETURNING id;
	`

	var id int
	err := r.DB.QueryRow(query, token, expirationDuration).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		if err == sql.ErrNoRows {
			return constants.TOKEN_KEEP_ALIVE_ERR
		}
		return constants.DB_OPERATION_ERR
	}

	return nil
}
