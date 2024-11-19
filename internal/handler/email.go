package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/email-notification-service/internal/service"
)

// RegisterEmailRoutes registers the email endpoints
func RegisterEmailRoutes(r *mux.Router) {
	r.HandleFunc("/send", SendEmail).Methods("POST")
}

// SendEmail handles the HTTP request to send an email
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var req service.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// go func() {
	// 	err := service.SendEmail(req)
	// 	if err != nil {
	// 		http.Error(w, "Failed to send email", http.StatusInternalServerError)
	// 		return
	// 	}
	// }()

	err := service.SendEmail(req)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}
