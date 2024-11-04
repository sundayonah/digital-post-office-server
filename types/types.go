package types

import "time"

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

// UpdateOrderRequest represents the request structure for updating an order.
type UpdateOrderRequest struct {
	Description string `json:"description"`
	Status      string `json:"status"` // Could be an enum or specific statuses you define
}

// RegisterUserRequest represents the request structure for registering a new user.
type RegisterUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
