package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pratikjethe/go-token-manager/config"
	"github.com/pratikjethe/go-token-manager/controllers"
	db "github.com/pratikjethe/go-token-manager/db"
	repositories "github.com/pratikjethe/go-token-manager/repository"
	"github.com/pratikjethe/go-token-manager/routes"
	"github.com/pratikjethe/go-token-manager/server"
	"github.com/pratikjethe/go-token-manager/services"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load db config
	dbConfig, err := db.GetConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Run Migrations
	err = db.DBMigrationUP(*dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Initiate DB connection
	dbConn, err := db.InitDB(dbConfig.DBUri)
	if err != nil {
		log.Fatalf("Error while connecting to DB : %v ", err.Error())
	}
	log.Print("Succesfully connected to DB!")

	// Initiate services, repositories, controllers
	tokenConfig := config.NewTokenConfig()
	tokenRepo := repositories.NewTokenRepository(dbConn)
	tokenService := services.NewTokenService(tokenRepo, tokenConfig)
	tokenController := controllers.NewTokenController(tokenService)

	routes.RegisterTokenRoutes(tokenController)

	// Start the HTTP server.
	serverConfig := server.NewServerConfig()

	server := server.NewServer(serverConfig)
	if err := server.Start(); err != nil {
		log.Fatal("Error while starting server", err.Error())
	}

	log.Print("Server Initiated!")

}
