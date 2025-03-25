package ports

import "github.com/1729asraful/microservices/order/internal/application/core/domain"

// DBPort defines the interface for the database operations on orders.
type DBPort interface {
	// Get retrieves an order by its unique ID
	Get(id string) (domain.Order, error)

	// Save stores the given order into the database
	Save(*domain.Order) error
}
