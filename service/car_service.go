package service

import (
	"database/sql"
	"errors"
	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/repository"
	"log"
)

type CarService struct {
	carRepo *repository.CarRepository
}

func NewCarService(carRepo *repository.CarRepository) *CarService {
	return &CarService{carRepo: carRepo}
}

func (s *CarService) CreateCar(regNums []string, db *sql.DB) ([]*model.Car, error) {
	log.Println("[DEBUG] Entering CreateCar service method with registration numbers:", regNums)

	// Check if registration numbers already exist in the database
	existingRegNums, err := repository.GetExistingRegNums(db)
	if err != nil {
		log.Println("[ERROR] Error retrieving existing registration numbers from the database:", err)
		return nil, err
	}

	// Debug log for existing registration numbers
	log.Println("[DEBUG] Existing registration numbers in the database:", existingRegNums)

	// Filter out registration numbers that already exist
	newRegNums := s.filterNewRegNums(regNums, existingRegNums)

	// Info log for new registration numbers
	log.Println("[INFO] New registration numbers to be inserted into the database:", newRegNums)

	// If all registration numbers already exist, return an error
	if len(newRegNums) == 0 {
		return nil, errors.New("all registration numbers already exist in the database")
	}

	// Make a request to the external API with new registration numbers
	cars, err := repository.CreateCar(newRegNums, db)
	if err != nil {
		log.Println("[ERROR] Error creating cars in the database:", err)
		return nil, err
	}

	// Info log for successfully inserted cars
	for _, car := range cars {
		log.Println("[INFO] Successfully inserted car with registration number:", car.RegNum)
	}

	// Debug log for function exit
	log.Println("[DEBUG] Exiting CreateCar service method")

	return cars, nil
}

func (s *CarService) GetCarByID(carID string) (*model.Car, error) {
	return repository.GetCarByID(carID)
}

func (s *CarService) UpdateCar(carID string, updatedCar model.Car) error {
	return repository.UpdateCar(carID, updatedCar)
}

func (s *CarService) DeleteCar(carID string) error {
	return repository.DeleteCar(carID)
}
func (s *CarService) filterNewRegNums(regNums, existingRegNums []string) []string {
	// Create a map to store existing registration numbers for efficient lookup
	existingRegNumsMap := make(map[string]bool)
	for _, regNum := range existingRegNums {
		existingRegNumsMap[regNum] = true
	}

	// Filter out registration numbers that already exist in the database
	var newRegNums []string
	for _, regNum := range regNums {
		if !existingRegNumsMap[regNum] {
			newRegNums = append(newRegNums, regNum)
		} else {
			log.Print("This reg number is already exist in database")
		}
	}
	return newRegNums
}
