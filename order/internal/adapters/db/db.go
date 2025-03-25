package db

import (
	"fmt"

	"github.com/1729asraful/microservices/order/internal/application/core/domain" // Adjust to your import path
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order represents the Order model to be persisted in the database
type Order struct {
	gorm.Model
	CustomerID int64       `gorm:"column:customer_id"`
	Status     string      `gorm:"column:status"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"` // Reference to OrderItem
}

// OrderItem represents an item in an order
type OrderItem struct {
	gorm.Model
	ProductCode string  `gorm:"column:product_code"`
	UnitPrice   float32 `gorm:"column:unit_price"`
	Quantity    int32   `gorm:"column:quantity"`
	OrderID     uint    `gorm:"column:order_id"` // Back reference to Order model
}

// Adapter struct holds a reference to the GORM DB instance
type Adapter struct {
	db *gorm.DB
}

// NewAdapter initializes a new Adapter with a DB connection using the provided DSN
func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	// Open the connection to the MySQL database
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// Migrate the database schema, creating the necessary tables
	err := db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	// Return the Adapter with the GORM DB instance
	return &Adapter{db: db}, nil
}

// Get retrieves an Order by its ID and maps it to the domain.Order model
func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	// Fetch the order entity from the database
	res := a.db.First(&orderEntity, id)

	// If there's an error (e.g., not found), return it
	if res.Error != nil {
		return domain.Order{}, res.Error
	}

	// Convert the OrderItem entities to domain models
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	// Map the OrderEntity to the domain model
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, nil
}

// Save stores a domain.Order in the database
func (a Adapter) Save(order *domain.Order) error {
	// Convert the domain OrderItems to OrderItem entities
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	// Convert the domain Order to an Order entity
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	// Save the order to the database
	res := a.db.Create(&orderModel)

	// If the save was successful, set the domain order's ID
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}

	return res.Error
}
