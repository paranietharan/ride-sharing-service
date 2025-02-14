package db

import (
	"fmt"
	"ride-sharing-service/pkg/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username string, password string, email string, phoneNo string) error {
	u := models.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phoneNo,
	}

	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(db *gorm.DB, userId int) (models.User, error) {
	var user models.User
	result := db.First(&user, userId)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetByPhoneNumber(db *gorm.DB, phoneNo string) (models.User, error) {
	var user models.User

	result := db.Where("phone = ?", phoneNo).First(&user)

	if result.Error != nil {
		fmt.Println("No user found with phone number:", phoneNo)
		return models.User{}, result.Error
	}

	return user, nil
}

func CreateNewRide(
	db *gorm.DB,
	customerPhone string,
	pickupLocation string,
	dropoffLocation string,
	vehicleType string,
	estimatedFare float64,
	companyID string,
	createdAt time.Time,
	users []models.User,
) (models.Ride, error) {

	ride := models.Ride{
		ID:              uuid.New().String(),
		CustomerPhone:   customerPhone,
		PickupLocation:  pickupLocation,
		DropoffLocation: dropoffLocation,
		VehicleType:     vehicleType,
		Status:          "PROCESSING",
		DriverAssigned:  false,
		EstimatedFare:   estimatedFare,
		CompanyID:       companyID,
		CreatedAt:       createdAt,
	}

	if err := db.Create(&ride).Error; err != nil {
		return models.Ride{}, err
	}

	for _, user := range users {
		if err := db.Model(&user).Association("Rides").Append(&ride); err != nil {
			return models.Ride{}, err
		}
	}

	return ride, nil
}

func MarkRideAsOngoing(db *gorm.DB, rideID string) error {
	return db.Model(&models.Ride{}).Where("id = ?", rideID).Updates(map[string]interface{}{
		"status":          "ONGOING",
		"driver_assigned": true,
	}).Error
}

func MarkRideAsFinished(db *gorm.DB, rideID string) (models.Ride, error) {
	err := db.Model(&models.Ride{}).Where("id = ?", rideID).Update("status", "COMPLETED").Error
	if err != nil {
		return models.Ride{}, err
	}

	r, _ := GetRideById(db, rideID)

	return r, nil
}

// list all rides by the company id
func ListAllRidesByCompanyID(db *gorm.DB, companyID string) ([]models.Ride, error) {
	var rides []models.Ride

	err := db.Where("company_id = ?", companyID).Find(&rides).Error
	if err != nil {
		return nil, err
	}

	return rides, nil
}

// get ride details by ride id
func GetRideById(db *gorm.DB, rideId string) (models.Ride, error) {
	var ride models.Ride
	err := db.Where("id = ?", rideId).Find(&ride).Error

	if err != nil {
		return models.Ride{}, err
	}

	return ride, nil
}

func GetRideByPhoneNumber(db *gorm.DB, phoneNo string) (models.Ride, error) {
	var ride models.Ride
	err := db.Where("customer_phone = ?", phoneNo).Find(&ride).Error

	if err != nil {
		return models.Ride{}, err
	}

	return ride, nil
}
