package routes

import (
	"net/http"

	"github.com/pratikjethe/go-token-manager/controllers"
)

func RegisterTokenRoutes(tokenController *controllers.TokenController) {
	http.HandleFunc("/tokens/create", tokenController.CreateTokenHandler)
	http.HandleFunc("/tokens/assign", tokenController.AssignToken)

}
