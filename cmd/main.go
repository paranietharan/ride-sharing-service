package main

import (
	"log"
	"ride-sharing-service/pkg/config"
	"ride-sharing-service/pkg/models"
)

func main() {
	config.Connect()
	defer config.Disconnect()

	// code to create the table if there is a first time start
	err := config.DB.AutoMigrate(&models.User{}, &models.Ride{}, &models.UserRides{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}

	log.Println("Tables created successfully!")
}
