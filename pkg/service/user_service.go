package service

import (
	"errors"
	"fmt"
	"ride-sharing-service/pkg/db"
	"ride-sharing-service/pkg/dto"
	"ride-sharing-service/pkg/models"
	"time"

	"gorm.io/gorm"
)

func CreateNewUser(database *gorm.DB, u dto.UserAccountCreateRequestDto) (dto.UserAccountCreateResponseDto, error) {
	existingUser, _ := db.GetByPhoneNumber(database, u.Phone)
	if existingUser.ID != 0 {
		return dto.UserAccountCreateResponseDto{}, errors.New("user already exists")
	}

	// Create new user
	err := db.CreateUser(database, u.Username, u.Password, u.Email, u.Phone)
	if err != nil {
		return dto.UserAccountCreateResponseDto{}, err
	}

	return dto.UserAccountCreateResponseDto{
		Username: u.Username,
		Email:    u.Email,
		Phone:    u.Phone,
		Message:  "User created successfully",
	}, nil
}

func RequestRide(database *gorm.DB, req dto.RideRequestDto) (dto.RideRequestResponseDto, error) {
	var users []models.User

	// Check if user exists
	u, err := db.GetByPhoneNumber(database, req.PhoneNo)
	if err != nil {
		return dto.RideRequestResponseDto{}, errors.New("user not found")
	}

	// Check if locations match hardcoded values
	fare, found := models.GetHardcodedRideLocation(req.PickUpLocation, req.DropOffLocation)
	if !found {
		return dto.RideRequestResponseDto{}, errors.New("no matching route found")
	}

	users = append(users, u)

	// Create a new ride
	r, err := db.CreateNewRide(
		database,
		u.Phone,
		req.PickUpLocation,
		req.DropOffLocation,
		req.VehicleType,
		fare,
		"DEFAULT",
		time.Now(),
		users,
	)
	if err != nil {
		return dto.RideRequestResponseDto{}, err
	}

	return toRideRequestResponseDto(r), nil
}

func FetchRideByRideId(database *gorm.DB, rideId string) (dto.RideDetailsDto, error) {
	r, e := db.GetRideById(database, rideId)

	if e != nil {
		return dto.RideDetailsDto{}, e
	}

	return toRideDetailsDto(r), nil
}

func FetchRideByRidePhoneNumber(database *gorm.DB, phoneNo string) (dto.RideDetailsDto, error) {
	r, e := db.GetRideByPhoneNumber(database, phoneNo)

	if e != nil {
		return dto.RideDetailsDto{}, e
	}

	return toRideDetailsDto(r), nil
}

func SubmitPayment(database *gorm.DB, req dto.SubmitPaymentRequestDto) (dto.SubmitPaymentResponseDto, error) {
	// validate ride that is in processing or ongoing
	r, e := db.GetRideById(database, req.RideId)

	if e != nil {
		return dto.SubmitPaymentResponseDto{}, e
	}

	// check the process
	if r.Status == "FINISHED" {
		return dto.SubmitPaymentResponseDto{}, fmt.Errorf("you are try to submit a payment to ended trip")
	}
	// then calculate the tip & fare amt and calcualte compare with the amount
	totalAmount := req.FareAmount + req.TipAmount

	tokenAmt, v := models.GetHardcodedPaymentToken(req.PaymentId)
	if !v {
		fmt.Println("Token not valid")
		return dto.SubmitPaymentResponseDto{}, fmt.Errorf("token is not valid")
	}

	if tokenAmt-totalAmount < 0 {
		fmt.Println("Amount mismatch")
		return dto.SubmitPaymentResponseDto{}, fmt.Errorf("total amount does not match the expected fare")
	}

	// then set the trip ststus as finished
	updatedRide, err := db.MarkRideAsFinished(database, req.RideId)
	if err != nil {
		return dto.SubmitPaymentResponseDto{}, err
	}

	res := dto.SubmitPaymentResponseDto{
		RideId:    updatedRide.ID,
		Status:    updatedRide.Status,
		TotalPaid: totalAmount,
	}

	return res, nil
}

func toRideRequestResponseDto(r models.Ride) dto.RideRequestResponseDto {
	rq := dto.RideRequestResponseDto{
		RideId:         r.ID,
		RideStatus:     r.Status,
		DriverAssigned: r.DriverAssigned,
		EstimatedFee:   r.EstimatedFare,
	}

	return rq
}

func toRideDetailsDto(ride models.Ride) dto.RideDetailsDto {
	return dto.RideDetailsDto{
		CustomerPhone:   ride.CustomerPhone,
		PickupLocation:  ride.PickupLocation,
		DropoffLocation: ride.DropoffLocation,
		VehicleType:     ride.VehicleType,
		Status:          ride.Status,
		DriverAssigned:  ride.DriverAssigned,
		EstimatedFare:   ride.EstimatedFare,
		CompanyID:       ride.CompanyID,
		CreatedAt:       ride.CreatedAt,
	}
}
