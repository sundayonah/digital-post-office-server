// handlers/order.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sundayonah/digital_post_office/ent"
	"github.com/sundayonah/digital_post_office/ent/order"
	"github.com/sundayonah/digital_post_office/ent/user"
	"github.com/sundayonah/digital_post_office/notification"
	"github.com/sundayonah/digital_post_office/types"
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

	var req types.CreateOrderRequest
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
	response := types.OrderResponse{
		ID:                 order.ID,
		TrackingNumber:     order.TrackingNumber,
		PackageDescription: order.PackageDescription,
		Status:             order.Status.String(),
		SafeCode:           order.SafeCode,
		CreatedAt:          order.CreatedAt,
		UpdatedAt:          order.UpdatedAt,
		Sender: types.UserResponse{
			FullName:  sender.FullName,
			Phone:     sender.Phone,
			Email:     sender.Email,
			CreatedAt: sender.CreatedAt,
		},
		Recipient: types.UserResponse{
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
	var response []types.OrderResponse
	for _, order := range orders {
		response = append(response, types.OrderResponse{
			ID:                 order.ID,
			TrackingNumber:     order.TrackingNumber,
			PackageDescription: order.PackageDescription,
			Status:             order.Status.String(),
			CreatedAt:          order.CreatedAt,
			UpdatedAt:          order.UpdatedAt,
			Sender: types.UserResponse{
				FullName:  order.Edges.Sender.FullName,
				Phone:     order.Edges.Sender.Phone,
				Email:     order.Edges.Sender.Email,
				CreatedAt: order.Edges.Sender.CreatedAt,
			},
			Recipient: types.UserResponse{
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

// GetOrderByID retrieves a specific order by its ID.
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id") // Assuming the ID is passed as a query parameter.

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
	}

	order, err := h.client.Order.Query().
		Where(order.ID(idInt)).
		WithSender().
		WithRecipient().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve order", http.StatusInternalServerError)
		return
	}

	response := types.OrderResponse{
		ID:                 order.ID,
		TrackingNumber:     order.TrackingNumber,
		PackageDescription: order.PackageDescription,
		Status:             order.Status.String(),
		CreatedAt:          order.CreatedAt,
		UpdatedAt:          order.UpdatedAt,
		Sender: types.UserResponse{
			FullName:  order.Edges.Sender.FullName,
			Phone:     order.Edges.Sender.Phone,
			Email:     order.Edges.Sender.Email,
			CreatedAt: order.Edges.Sender.CreatedAt,
		},
		Recipient: types.UserResponse{
			FullName:  order.Edges.Recipient.FullName,
			Phone:     order.Edges.Recipient.Phone,
			Email:     order.Edges.Recipient.Email,
			CreatedAt: order.Edges.Recipient.CreatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateOrder updates an existing order.
// func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	idStr := r.URL.Query().Get("id")

// 	// convert ID from string to int (assuming ID is an int)
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid order ID", http.StatusBadRequest)
// 		return
// 	}

// 	var req types.UpdateOrderRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	order, err := h.client.Order.Get(ctx, id)
// 	if err != nil {
// 		if ent.IsNotFound(err) {
// 			http.Error(w, "Order not found", http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, "Failed to retrieve order", http.StatusInternalServerError)
// 		return
// 	}

// 	// Assuming you have a function like StatusFromString that converts string to types.OrderStatus
// 	status, err := types.StatusFromString(req.Status)
// 	if err != nil {
// 		http.Error(w, "Invalid status", http.StatusBadRequest)
// 		return
// 	}

// 	// Convert types.OrderStatus to the appropriate enum value
// 	var entStatus string
// 	switch status {
// 	case types.StatusPending:
// 		entStatus = order.StatusPending
// 	case types.StatusDelivered:
// 		entStatus = order.StatusDelivered
// 	case types.StatusCancelled:
// 		entStatus = order.StatusCancelled
// 	default:
// 		http.Error(w, "Invalid status", http.StatusBadRequest)
// 		return
// 	}

// 	// Update the order details
// 	order, err = h.client.Order.UpdateOne(order).
// 		SetPackageDescription(req.Description).
// 		SetStatus(entStatus).
// 		Save(ctx)

// 	response := types.OrderResponse{
// 		ID:                 order.ID,
// 		TrackingNumber:     order.TrackingNumber,
// 		PackageDescription: order.PackageDescription,
// 		Status:             order.Status.String(),
// 		CreatedAt:          order.CreatedAt,
// 		UpdatedAt:          order.UpdatedAt,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// DeleteOrder removes an order by its ID.
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
	}

	err = h.client.Order.DeleteOneID(idInt).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content for successful deletion.
}

// RegisterUser creates a new user in the system.
func (h *OrderHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req types.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.client.User.Create().
		SetFullName(req.FullName).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		Save(ctx)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	response := types.UserResponse{
		// ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserOrders retrieves all orders for a specific user.
func (h *OrderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("user_id")

	userInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
	}

	// Fetch orders where the user is either the sender or recipient
	orders, err := h.client.Order.Query().
		Where(order.HasSenderWith(user.ID(userInt)), (order.HasRecipientWith(user.ID(userInt)))).
		WithRecipient(). // This will load the recipient details
		WithSender().    // If you also want to load the sender details
		All(ctx)

	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	var response []types.OrderResponse
	for _, order := range orders {
		response = append(response, types.OrderResponse{
			ID:                 order.ID,
			TrackingNumber:     order.TrackingNumber,
			PackageDescription: order.PackageDescription,
			Status:             order.Status.String(),
			CreatedAt:          order.CreatedAt,
			UpdatedAt:          order.UpdatedAt,
			Sender: types.UserResponse{
				FullName:  order.Edges.Sender.FullName,
				Phone:     order.Edges.Sender.Phone,
				Email:     order.Edges.Sender.Email,
				CreatedAt: order.Edges.Sender.CreatedAt,
			},
			Recipient: types.UserResponse{
				FullName:  order.Edges.Recipient.FullName,
				Phone:     order.Edges.Recipient.Phone,
				Email:     order.Edges.Recipient.Email,
				CreatedAt: order.Edges.Recipient.CreatedAt,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
