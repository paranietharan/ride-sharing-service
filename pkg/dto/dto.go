package dto

// for user creation endpoint
type UserAccountCreateRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserAccountCreateResponseDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Message  string
}

type RideRequestDto struct {
	PhoneNo         string `json:"phoneNo"`
	PickUpLocation  string `json:"pickUpLocation"`
	DropOffLocation string `json:"dropOffLocation"`
	VehicleType     string `json:"vehicleType"`
}

type RideRequestResponseDto struct {
	RideId         string
	RideStatus     string
	DriverAssigned bool
	EstimatedFee   float64
}

type RideDto struct {
	RideId         int
	Status         string
	DriverAssigned bool
	EstimatedFare  float64
}
