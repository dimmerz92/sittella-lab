package main

import (
	"net/http"
	"os"

	"sittellalab.com.au/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/api/contact", api.EnquiryHandler)
	http.HandleFunc("/api/chat", api.GPTChatHandler)

	http.ListenAndServe(":"+port, nil)
}
