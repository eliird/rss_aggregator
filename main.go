package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handler_readiness)
	mux.HandleFunc("GET /v1/err", handler_error)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server started on port http://localhost:%s", port)
	log.Fatal(srv.ListenAndServe())

}
