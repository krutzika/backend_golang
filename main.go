package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/krutzika/backend_golang/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	request := router.SetupRouter()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", request))
}
