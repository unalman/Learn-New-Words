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
	router := server.SetupRouter()
	router = GetWords(server, router)
	router = GetRandomWords(server, router)
	router = GetWord(server, router)
	router = PostWord(server, router)
	router = UpdateWord(server, router)
	router = DeleteWord(server, router)

	router.Run(server.listenAddr)
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
