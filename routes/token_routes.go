package routes

import (
	"net/http"

	"github.com/pratikjethe/go-token-manager/controllers"
)

func RegisterTokenRoutes(tokenController *controllers.TokenController) {
	http.HandleFunc("/tokens/create", tokenController.CreateTokenHandler)
	http.HandleFunc("/tokens/assign", tokenController.AssignTokenHandler)
	http.HandleFunc("/tokens/delete", tokenController.DeleteTokenHandler)
	http.HandleFunc("/tokens/unblock", tokenController.UnblockTokenHandler)
	http.HandleFunc("/tokens/keep-alive", tokenController.KeepAliveTokenHandler)

}
