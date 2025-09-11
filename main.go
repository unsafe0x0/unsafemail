package main

import (
	"log"
	"net/http"
	"unsafemail/api"
	"unsafemail/config"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/send-email", api.EmailHandler)

	loggedMux := loggingMiddleware(mux)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}
