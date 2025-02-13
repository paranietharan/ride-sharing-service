package dto

type RideDto struct {
	RideId         int
	Status         string
	DriverAssigned bool
	EstimatedFare  float64
}
