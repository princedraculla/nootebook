package main

import (
	"github.com/joho/godotenv"
	"log"
	"nootebook.com/internal/gateway/http"
)

func init() {
	err := godotenv.Load("/home/amir/GitHub/phoonebook/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
func main() {
	http.ServerInit()
}
