package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	getEnv()
	store, err := CreatePostgreSqlConnection()
	if err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), store)
	server.Run()
}

// this method load environment variables
func getEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" || env == "dev" {
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Panic(err)
		}
	} else {
		err := godotenv.Load(".env.prod")
		if err != nil {
			log.Panic(err)
		}
	}
}
