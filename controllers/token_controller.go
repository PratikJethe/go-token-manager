package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/pratikjethe/go-token-manager/services"
)

type TokenController struct {
	tokenService *services.TokenService
}

func NewTokenController(service *services.TokenService) *TokenController {
	return &TokenController{tokenService: service}
}

func (controller *TokenController) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {

	tokenSizeStr := os.Getenv("TOKEN_SIZE")
	tokenSize, err := strconv.Atoi(tokenSizeStr)
	if err != nil || tokenSize <= 0 {
		tokenSize = 16
	}

	token, err := controller.tokenService.CreateToken(tokenSize)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
