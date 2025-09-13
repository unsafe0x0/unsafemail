package main

import (
	"log"
	"net/http"
	"os"
	"unsafemail/api"
	"unsafemail/config"

	"github.com/joho/godotenv"
)

var corsOrigin string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	corsOrigin = os.Getenv("CORS_ORIGIN")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/send-email", api.EmailHandler)

	handler := loggingMiddleware(corsMiddleware(mux))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
