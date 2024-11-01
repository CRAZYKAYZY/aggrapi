package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChileKasoka/mis/cmd/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the PORT environment variable
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Initialize the application
	application, err := app.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	fmt.Println("Initialized successfully")

	// Start the server
	err = application.Server.Start(":" + portString)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	fmt.Println("Server started on port:", portString)
}
