package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/pratikjethe/go-token-manager/constants"
	"github.com/pratikjethe/go-token-manager/services"
)

type TokenController struct {
	tokenService *services.TokenService
}

func NewTokenController(service *services.TokenService) *TokenController {
	return &TokenController{tokenService: service}
}

func (controller *TokenController) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid http method"})
		return
	}

	tokenSizeStr := os.Getenv("TOKEN_SIZE")
	tokenSize, err := strconv.Atoi(tokenSizeStr)
	if err != nil || tokenSize <= 0 {
		tokenSize = 16
	}

	token, err := controller.tokenService.CreateToken()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not create token"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func (c *TokenController) AssignToken(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid http method"})
		return
	}

	token, err := c.tokenService.AssignToken()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err.Error() == constants.ERR_NO_TOKENS.Error() {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "No available tokens to assign"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	response := map[string]string{"token": token.Token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *TokenController) DeleteTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get the token from the URL parameters

	if r.Method != http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid http method"})
		return
	}
	token := r.URL.Query().Get("token")
	if token == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token string is required"})

		return
	}

	// Call the service to delete the token
	err := c.tokenService.DeleteToken(token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		if err.Error() == constants.ERR_TOKEN_ALREADY_DELETED.Error() {
			json.NewEncoder(w).Encode(map[string]string{"error": "Token Not found or already deleted"})
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while deleting token"})
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Token deleted"}) // 204 No Content
}
