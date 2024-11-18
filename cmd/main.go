package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/email-notification-service/internal/handler"
)

func main() {
	r := mux.NewRouter()

	handler.RegisterEmailRoutes(r)

	// Start the HTTP server
	log.Println("Email Notification Service running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
