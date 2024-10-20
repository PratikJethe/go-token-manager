package services

import (
	"time"

	"github.com/pratikjethe/go-token-manager/models"
	repositories "github.com/pratikjethe/go-token-manager/repository"
	"github.com/pratikjethe/go-token-manager/utils"
)

type TokenService struct {
	tokenRepo *repositories.TokenRepository
}

func NewTokenService(repo *repositories.TokenRepository) *TokenService {
	return &TokenService{tokenRepo: repo}
}

func (service *TokenService) CreateToken(tokenSize int) (*models.Token, error) {

	t, err := utils.GenerateRandomToken(tokenSize)

	if err != nil {
		return nil, err
	}
	token := &models.Token{
		Token:              t,
		LastActivationTime: time.Now().UTC(),
		IsDeleted:          false,
	}

	err = service.tokenRepo.CreateToken(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
