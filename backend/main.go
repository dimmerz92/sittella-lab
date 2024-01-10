package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"sittellalab.com.au/internal/api"
	"sittellalab.com.au/internal/middleware"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/contact", api.EnquiryHandler)
	mux.HandleFunc("/api/chat", api.GPTChatHandler)

	server := middleware.NewValidator(mux)

	http.ListenAndServe(":"+port, server)
}
