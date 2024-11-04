// handlers/order.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sundayonah/digital_post_office/ent"
	"github.com/sundayonah/digital_post_office/ent/user"
	"github.com/sundayonah/digital_post_office/notification"
)

// OrderHandler struct manages order operations.
type OrderHandler struct {
	client              *ent.Client
	notificationService *notification.NotificationService
}

// NewOrderHandler creates a new instance of OrderHandler.
func NewOrderHandler(client *ent.Client, ns *notification.NotificationService) *OrderHandler {
	return &OrderHandler{
		client:              client,
		notificationService: ns,
	}
}

// CreateOrder handles the creation of a new order.
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Start transaction
	tx, err := h.client.Tx(ctx)
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// Find or create the sender
	sender, err := tx.User.Query().Where(user.Phone(req.SenderPhone)).Only(ctx)
	if ent.IsNotFound(err) {
		sender, err = tx.User.Create().
			SetFullName(req.SenderFullName).
			SetPhone(req.SenderPhone).
			SetEmail(req.SenderEmail).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create sender", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		tx.Rollback()
		http.Error(w, "Error finding sender", http.StatusInternalServerError)
		return
	}

	// Find or create the recipient
	recipient, err := tx.User.Query().Where(user.Phone(req.RecipientPhone)).Only(ctx)
	if ent.IsNotFound(err) {
		recipient, err = tx.User.Create().
			SetFullName(req.RecipientFullName).
			SetPhone(req.RecipientPhone).
			SetEmail(req.RecipientEmail).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create recipient", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		tx.Rollback()
		http.Error(w, "Error finding recipient", http.StatusInternalServerError)
		return
	}

	// Create the order
	order, err := tx.Order.Create().
		SetSender(sender).
		SetRecipient(recipient).
		SetPackageDescription(req.Description).
		SetTrackingNumber(req.TrackingNumber).
		SetStatus("pending").
		Save(ctx)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Generate safe_code
	safeCode := notification.GenerateSafeCode() // Use your safe code generation function
	order, err = tx.Order.UpdateOneID(order.ID).
		SetSafeCode(safeCode).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update safe code", http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	// Asynchronously notify the sender and recipient
	go func() {
		h.notificationService.NotifyNewOrder(ctx, order, sender, recipient)
	}()

	// Format and send response including safe_code
	response := OrderResponse{
		ID:                 order.ID,
		TrackingNumber:     order.TrackingNumber,
		PackageDescription: order.PackageDescription,
		Status:             order.Status.String(),
		SafeCode:           order.SafeCode,
		CreatedAt:          order.CreatedAt,
		UpdatedAt:          order.UpdatedAt,
		Sender: UserResponse{
			FullName:  sender.FullName,
			Phone:     sender.Phone,
			Email:     sender.Email,
			CreatedAt: sender.CreatedAt,
		},
		Recipient: UserResponse{
			FullName:  recipient.FullName,
			Phone:     recipient.Phone,
			Email:     recipient.Email,
			CreatedAt: recipient.CreatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllOrders retrieves all orders with sender and recipient details.
func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Query all orders with sender and recipient edges
	orders, err := h.client.Order.
		Query().
		WithSender().
		WithRecipient().
		All(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	// Map orders to custom response
	var response []OrderResponse
	for _, order := range orders {
		response = append(response, OrderResponse{
			ID:                 order.ID,
			TrackingNumber:     order.TrackingNumber,
			PackageDescription: order.PackageDescription,
			Status:             order.Status.String(),
			CreatedAt:          order.CreatedAt,
			UpdatedAt:          order.UpdatedAt,
			Sender: UserResponse{
				FullName:  order.Edges.Sender.FullName,
				Phone:     order.Edges.Sender.Phone,
				Email:     order.Edges.Sender.Email,
				CreatedAt: order.Edges.Sender.CreatedAt,
			},
			Recipient: UserResponse{
				FullName:  order.Edges.Recipient.FullName,
				Phone:     order.Edges.Recipient.Phone,
				Email:     order.Edges.Recipient.Email,
				CreatedAt: order.Edges.Recipient.CreatedAt,
			},
		})
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode orders to JSON", http.StatusInternalServerError)
		return
	}
}

// CreateOrderRequest represents the JSON request structure for creating an order.
type CreateOrderRequest struct {
	TrackingNumber    string `json:"tracking_number"`
	SenderFullName    string `json:"sender_full_name"`
	SenderPhone       string `json:"sender_phone"`
	SenderEmail       string `json:"sender_email"`
	RecipientFullName string `json:"recipient_full_name"`
	RecipientPhone    string `json:"recipient_phone"`
	RecipientEmail    string `json:"recipient_email"`
	Description       string `json:"description"`
}

// OrderResponse represents the JSON response structure for orders.
type OrderResponse struct {
	ID                 int          `json:"id"`
	TrackingNumber     string       `json:"tracking_number"`
	PackageDescription string       `json:"package_description"`
	Status             string       `json:"status"`
	SafeCode           string       `json:"safe_code"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
	Sender             UserResponse `json:"sender"`
	Recipient          UserResponse `json:"recipient"`
}

// UserResponse represents the JSON response structure for a user.
type UserResponse struct {
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
