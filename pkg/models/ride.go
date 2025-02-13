package models

import "time"

type Ride struct {
	RideID          int       `json:"rideId"`
	CustomerPhone   string    `json:"customerPhone"`
	PickupLocation  string    `json:"pickupLocation"`
	DropoffLocation string    `json:"dropoffLocation"`
	VehicleType     string    `json:"vehicleType"`
	Status          string    `json:"status"`
	DriverAssigned  bool      `json:"driverAssigned"`
	EstimatedFare   float64   `json:"estimatedFare"`
	CompanyID       string    `json:"companyId"`
	CreatedAt       time.Time `json:"createdAt"`
	Users           []User    `json:"users"`
}
