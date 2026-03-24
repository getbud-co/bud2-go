package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dsbraz/bud2/backend/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := handler.NewRouter()

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
