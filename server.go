package main

import (
	"fmt"
	"os"

	configs "go-mvc-server/configs"
	Routes "go-mvc-server/routes"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup logrus
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})

	configs.ConnectDB()
}

func main() {
	fmt.Println("Let's begin Boom Boom !")

	// Setup and run the router
	r := Routes.SetupRouter()
	fmt.Printf("Listening to port %s", "8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
