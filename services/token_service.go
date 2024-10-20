package services

import (
	"github.com/pratikjethe/go-token-manager/config"
	"github.com/pratikjethe/go-token-manager/models"
	repositories "github.com/pratikjethe/go-token-manager/repository"
	"github.com/pratikjethe/go-token-manager/utils"
)

type TokenService struct {
	tokenRepo   *repositories.TokenRepository
	tokenConfig *config.TokenConfig
}

func NewTokenService(repo *repositories.TokenRepository, c *config.TokenConfig) *TokenService {
	return &TokenService{tokenRepo: repo, tokenConfig: c}
}

func (s *TokenService) CreateToken() (*models.Token, error) {

	t, err := utils.GenerateRandomToken(s.tokenConfig.TokenLength)

	if err != nil {
		return nil, err
	}
	token := &models.Token{
		Token: t,
	}

	err = s.tokenRepo.CreateToken(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *TokenService) AssignToken() (*models.Token, error) {
	tx, err := s.tokenRepo.DB.Begin()
	if err != nil {
		return nil, err
	}

	token, err := s.tokenRepo.GetAvailableToken(tx, s.tokenConfig.TokenActiveDuration, s.tokenConfig.TokenExpireDuration)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = s.tokenRepo.AssignToken(tx, token)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) DeleteToken(token string) error {
	return s.tokenRepo.DeleteToken(token)
}

func (s *TokenService) UnblockToken(token string) error {
	// Use the repository to update the token status
	err := s.tokenRepo.UnblockToken(token, s.tokenConfig.TokenActiveDuration)
	if err != nil {
		return err
	}
	return nil
}

func (s *TokenService) KeepAliveToken(token string) error {
	err := s.tokenRepo.KeepAliveToken(token, s.tokenConfig.TokenExpireDuration)
	if err != nil {
		return err
	}
	return nil
}
