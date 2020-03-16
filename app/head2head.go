package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//app to receive head 2 head of desired player vs the current top 20 players of
//desired pr
func getAPIKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}
	return os.Getenv(key)
}
func main() {
	//set .env variable API_KEY for ur api key to use
	apiKey := getAPIKey("API_KEY")
	println(apiKey)
}
