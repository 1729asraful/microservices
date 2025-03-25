package main

import (
	"log"

	"github.com/1729asraful/microservices/order/config"
	"github.com/1729asraful/microservices/order/internal/adapters/db"
	"github.com/1729asraful/microservices/order/internal/adapters/grpc"
	"github.com/1729asraful/microservices/order/internal/application/core/api"
)

func main() {
	// Fetch environment variables and initialize adapters
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// Create an instance of the core application
	application := api.NewApplication(dbAdapter)

	// Create the gRPC adapter and start the server
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
