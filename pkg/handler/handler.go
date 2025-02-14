package handler

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/pkg/dto"
	"ride-sharing-service/pkg/service"
	"ride-sharing-service/pkg/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateNewUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userReq dto.UserAccountCreateRequestDto

		err := json.NewDecoder(r.Body).Decode(&userReq)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		res, err := service.CreateNewUser(db, userReq)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusConflict, err.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, res)
	}
}

func RequestRide(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rideReq dto.RideRequestDto

		err := json.NewDecoder(r.Body).Decode(&rideReq)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Request Body")
			return
		}

		res, err := service.RequestRide(db, rideReq)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		utils.WriteResponse(w, http.StatusCreated, res)
	}
}

func FetchRideByRideId(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		rideId := vars["rideId"]

		rideDetails, err := service.FetchRideByRideId(db, rideId)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Ride not found")
			return
		}

		utils.WriteResponse(w, http.StatusOK, rideDetails)
	}
}

func FetchRideByRideByPhoneNumber(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		phoneNo := vars["phoneNo"]

		rideDetails, err := service.FetchRideByRidePhoneNumber(db, phoneNo)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "Ride not found")
			return
		}

		utils.WriteResponse(w, http.StatusOK, rideDetails)
	}
}
