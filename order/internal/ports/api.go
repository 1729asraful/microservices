package ports

import "github.com/1729asraful/microservices/order/internal/application/core/domain"

// APIPort defines the interface for the API layer to handle orders.
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
