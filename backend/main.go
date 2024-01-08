package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"sittellalab.com.au/internal/api"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/api/contact", api.EnquiryHandler)
	http.HandleFunc("/api/chat", api.GPTChatHandler)

	http.ListenAndServe(":"+port, nil)
}
