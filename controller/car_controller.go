package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/service"
)

type CarController struct {
	service *service.CarService
	db      *sql.DB
}

func NewCarController(carService *service.CarService, db *sql.DB) *CarController {
	return &CarController{
		service: carService,
		db:      db,
	}
}

func (c *CarController) CreateCarHandler(w http.ResponseWriter, r *http.Request) {
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

func (c *CarController) GetCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "Car ID is required", http.StatusBadRequest)
		return
	}

	car, err := c.service.GetCarByID(carID)
	if err != nil {
		http.Error(w, "Failed to retrieve car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func (c *CarController) UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
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
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CarController) DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")
	if carID == "" {
		http.Error(w, "Car ID is required", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteCar(carID); err != nil {
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getCarInfo(regNum string) {
	// Create a GET request to the external API endpoint
	url := fmt.Sprintf("https://external-api.com/info?regNum=%s", regNum)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to call external API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Decode the response body
	var carInfo model.Car
	if err := json.NewDecoder(resp.Body).Decode(&carInfo); err != nil {
		fmt.Printf("Failed to decode API response: %v\n", err)
		return
	}

	// Process the car information as needed
	fmt.Println("Received car info from external API:", carInfo)
}

func defineMockCarHandler(regNum string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the response body (imaginary data)
		responseData := map[string]interface{}{
			"regNum": regNum,
			"mark":   "Lada",
			"model":  "Vesta",
			"year":   2002,
			"owner": map[string]string{
				"name":       "John",
				"surname":    "Doe",
				"patronymic": "Smith",
			},
		}

		// Set the response status code to 200 (OK)
		w.WriteHeader(http.StatusOK)

		// Marshal the response data to JSON format
		responseJSON, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Write the JSON response to the response writer
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	})
}
