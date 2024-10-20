package main

import (
	"log"

	"github.com/joho/godotenv"
	db "github.com/pratikjethe/go-token-manager/db"
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
	_, err = db.InitDB(dbConfig.DBUri)
	if err != nil {
		log.Fatalf("Error while connecting to DB : %v ", err.Error())
	}

	log.Print("Succesfully connected to DB!")

}
