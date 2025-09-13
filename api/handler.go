package api

import (
	"encoding/json"
	"log"
	"net/http"
	"unsafemail/email"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("invalid method %s on %s", r.Method, r.URL.Path)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("failed to decode request: %v", err)
		return
	}

	if err := email.Send(req.To, req.Subject, req.Body); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		log.Printf("failed to send email to %s: %v", req.To, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
	log.Printf("email sent to %s successfully", req.To)
}
