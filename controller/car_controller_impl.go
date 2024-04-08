package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/service"
	"log"
	"net/http"
	"strconv"
)

// CarControllerImpl implements CarController interface
type CarControllerImpl struct {
	service service.CarService
	db      *sql.DB
}

// NewCarController creates a new instance of CarControllerImpl
func NewCarController(carService service.CarService, db *sql.DB) *CarControllerImpl {
	return &CarControllerImpl{
		service: carService,
		db:      db,
	}
}

func (c *CarControllerImpl) CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		RegNums []string `json:"regNums"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Iterate over each registration number and process them individually
	var createdCars []*model.Car
	for _, regNum := range requestData.RegNums {
		// Call the service method with each registration number and the database pointer
		createdCar, err := c.service.CreateCar([]string{regNum}, c.db)
		if err != nil {
			http.Error(w, "Failed to create car", http.StatusInternalServerError)
			return
		}
		createdCars = append(createdCars, createdCar...)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCars)
}

func (c *CarControllerImpl) GetCarsFilteredHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the filtering criteria from the query parameters
	criteria := r.URL.Query().Get("criteria")
	if criteria == "" {
		http.Error(w, "Filtering criteria are required", http.StatusBadRequest)
		return
	}

	// Extract the limit and offset for pagination
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	// Convert limit and offset to integers
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
		return
	}

	// Call the service method to get cars based on filtering criteria
	cars, owners, err := c.service.GetCarsFiltered(criteria, limitInt, offsetInt)
	if err != nil {
		http.Error(w, "Failed to retrieve cars", http.StatusInternalServerError)
		return
	}

	// Combine cars and owners into a single response struct
	var response []struct {
		Car   *model.Car    `json:"car"`
		Owner *model.People `json:"owner"`
	}

	// Combine cars and owners into the response
	for i, car := range cars {
		response = append(response, struct {
			Car   *model.Car    `json:"car"`
			Owner *model.People `json:"owner"`
		}{
			Car:   car,
			Owner: owners[i],
		})
	}

	// Set response headers and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *CarControllerImpl) UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "Car ID is required", http.StatusBadRequest)
		return
	}

	var updatedCar model.Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := c.service.UpdateCar(carID, updatedCar); err != nil {
		http.Error(w, "Failed to update car", http.StatusInternalServerError)
		return
	} else {
		log.Println("[INFO] Successfully updated car with car id:", carID)
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CarControllerImpl) DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "Car ID is required", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteCar(carID); err != nil {
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	} else {
		log.Println("[INFO] Successfully deleted car with car id:", carID)
	}

	w.WriteHeader(http.StatusOK)
}
