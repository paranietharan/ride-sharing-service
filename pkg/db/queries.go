package db

import (
	"ride-sharing-service/pkg/models"
	"time"

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

	result := db.First(&user, user.ID)

	if result.Error != nil {
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

	err := db.Create(&ride).Error
	if err != nil {
		return models.Ride{}, err
	}

	for _, user := range users {
		err := db.Model(&user).Association("Rides").Append(&ride)
		if err != nil {
			return models.Ride{}, err
		}
	}

	return ride, nil
}

func MarkRideAsOngoing(db *gorm.DB, rideID string) error {
	return db.Model(&models.Ride{}).Where("ride_id = ?", rideID).Updates(map[string]interface{}{
		"status":          "ONGOING",
		"driver_assigned": true,
	}).Error
}

func MarkRideAsFinished(db *gorm.DB, rideID string) error {
	return db.Model(&models.Ride{}).Where("ride_id = ?", rideID).Update("status", "COMPLETED").Error
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
