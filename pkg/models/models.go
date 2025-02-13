package models

import "time"

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Rides    []Ride `json:"rides" gorm:"many2many:user_rides;"`
}

type Ride struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	CustomerPhone   string    `json:"customerPhone"`
	PickupLocation  string    `json:"pickupLocation"`
	DropoffLocation string    `json:"dropoffLocation"`
	VehicleType     string    `json:"vehicleType"`
	Status          string    `json:"status"` //processing, ongoing, finished
	DriverAssigned  bool      `json:"driverAssigned"`
	EstimatedFare   float64   `json:"estimatedFare"`
	CompanyID       string    `json:"companyId"`
	CreatedAt       time.Time `json:"createdAt"`
	Users           []User    `json:"users" gorm:"many2many:user_rides;"`
}

type UserRides struct {
	UserID int `gorm:"primaryKey"`
	RideID int `gorm:"primaryKey"`
}
