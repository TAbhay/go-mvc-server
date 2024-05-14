package main

import (
	"fmt"
	Routes "go-mvc-server/routes"
	"os"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func init() {
	// load environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup logrus
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	fmt.Println("Let's begin Boom Boom !")
	r := Routes.SetupRouter()
	fmt.Printf("Listening to port %s", "8080")
	r.Run(":8080")

}
