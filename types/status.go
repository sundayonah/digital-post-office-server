package types

import (
	"fmt"
)

// Define the OrderStatus type
type OrderStatus string

// const (
// 	OrderStatusPending   OrderStatus = "pending"
// 	OrderStatusDelivered OrderStatus = "delivered"
// 	OrderStatusCancelled OrderStatus = "cancelled"
// )

const (
	StatusPending   = "pending"
	StatusDelivered = "delivered"
	StatusCancelled = "cancelled"
)

// StatusFromString converts a string to an OrderStatus type
func StatusFromString(status string) (OrderStatus, error) {
	switch status {
	case string(StatusPending):
		return StatusPending, nil
	case string(StatusDelivered):
		return StatusDelivered, nil
	case string(StatusCancelled):
		return StatusCancelled, nil
	default:
		return "", fmt.Errorf("invalid status: %s", status)
	}
}
