package api

import (
	"github.com/1729asraful/microservices/order/internal/application/core/domain"
	"github.com/1729asraful/microservices/order/internal/ports"
)

// Application represents the API for managing orders.
type Application struct {
	db ports.DBPort
}

// NewApplication initializes the Application with the given DB port.
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db, // DBPort is passed during the app's initialization.
	}
}

// PlaceOrder places an order and saves it to the database.
func (a *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// Order is saved through the DB port.
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
