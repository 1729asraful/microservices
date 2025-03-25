package domain

import "time"

// OrderItem represents an item in an order
type OrderItem struct {
	ProductCode string  `json:"product_code"` // Unique code of the product
	UnitPrice   float32 `json:"unit_price"`   // Price of a single product
	Quantity    int32   `json:"quantity"`     // Count of the product
}

// Order represents an order with order items and status
type Order struct {
	ID         int64       `json:"id"`          // Unique identifier of the order
	CustomerID int64       `json:"customer_id"` // Customer who owns the order
	Status     string      `json:"status"`      // Status of the order
	OrderItems []OrderItem `json:"order_items"` // List of items purchased in an order
	CreatedAt  int64       `json:"created_at"`  // Order creation time
}

// NewOrder creates a new order
func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
