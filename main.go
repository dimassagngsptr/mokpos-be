package main

import (
	"log"
	"os"
	"transaction-be/src"

	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return "0.0.0.0:" + port
}

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	app := src.App()

	if err := app.Listen(getPort()); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
