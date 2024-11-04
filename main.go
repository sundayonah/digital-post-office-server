// main.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/sundayonah/digital_post_office/database"
	"github.com/sundayonah/digital_post_office/handlers"
	"github.com/sundayonah/digital_post_office/middleware"
	"github.com/sundayonah/digital_post_office/notification"
)

func main() {
	// Create a background context
	ctx := context.Background()

	// Initialize the database
	client, err := database.InitDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer client.Close()

	// Initialize notification
	smsProvider := &notification.TwilioSMSProvider{} // Implement your SMS provider
	notificationService := notification.NewNotificationService(client, smsProvider)

	// Initialize handlers
	orderHandler := handlers.NewOrderHandler(client, notificationService)

	// Setup router
	r := mux.NewRouter()

	// Enable CORS for all routes
	r.Use(middleware.CorsMiddleware)

	// Define routes
	r.HandleFunc("/api/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/api/orders", orderHandler.GetAllOrders).Methods("GET")
	r.HandleFunc("/api/orders/{id}", orderHandler.GetOrderByID).Methods("GET")
	r.HandleFunc("/api/users/orders", orderHandler.GetUserOrders).Methods("GET")

	// r.HandleFunc("/api/orders/{id}", orderHandler.UpdateOrder).Methods("PUT")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
