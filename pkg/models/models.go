package models

import "time"

type Ride struct {
	RideID          string    `json:"rideId"`
	CustomerPhone   string    `json:"customerPhone"`
	PickupLocation  string    `json:"pickupLocation"`
	DropoffLocation string    `json:"dropoffLocation"`
	VehicleType     string    `json:"vehicleType"`
	Status          string    `json:"status"` // "pending", "completed", "cancelled"
	DriverAssigned  bool      `json:"driverAssigned"`
	EstimatedFare   float64   `json:"estimatedFare"`
	CompanyID       string    `json:"companyId"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Payment struct {
	RideID      string    `json:"rideId"`
	PaymentID   string    `json:"paymentId"`
	FareAmount  float64   `json:"fareAmount"`
	TipAmount   float64   `json:"tipAmount"`
	TotalAmount float64   `json:"totalPaid"`
	PaymentDate time.Time `json:"paymentDate"`
}
