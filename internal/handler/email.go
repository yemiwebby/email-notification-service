package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/email-notification-service/internal/service"
)

func RegisterEmailRoutes(r *mux.Router) {
	r.HandleFunc("/send", SendEmail).Methods("POST")
}

// func SendEmail(w http.ResponseWriter, r *http.Request) {
// 	start := time.Now()

// 	var req service.EmailRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	go func() {
// 		err := service.SendEmail(req)
// 		if err != nil {
// 			http.Error(w, "Failed to send email", http.StatusInternalServerError)
// 			return
// 		}
// 	}()

// 	log.Printf("Email processing completed in %v\n", time.Since(start))
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
// }


func SendEmail(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var req service.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Synchronous email sending (without concurrency)
	err := service.SendEmail(req)
	if err != nil {
		log.Printf("Failed to send email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Printf("Email processing completed in %v\n", time.Since(start))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}