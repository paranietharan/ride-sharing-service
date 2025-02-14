package main

import (
	"fmt"
	"log"
	"net/http"
	"ride-sharing-service/pkg/config"
	"ride-sharing-service/pkg/models"
	"ride-sharing-service/pkg/router"
)

func main() {
	db := config.Connect()
	defer config.Disconnect()

	err := db.AutoMigrate(&models.User{}, &models.Ride{}, &models.UserRides{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}
	log.Println("Database tables migrated successfully!")

	r := router.InitializeRoutes(db)

	port := ":8080"
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
