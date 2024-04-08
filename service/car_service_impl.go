package service

import (
	"database/sql"
	"errors"
	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/repository"
	"log"
)

// CarServiceImpl implements CarService interface
type CarServiceImpl struct {
	carRepo repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) *CarServiceImpl {
	return &CarServiceImpl{carRepo: carRepo}
}

func (s *CarServiceImpl) CreateCar(regNums []string, db *sql.DB) ([]*model.Car, error) {
	log.Println("[DEBUG] Entering CreateCar service method with registration numbers:", regNums)

	// Check if registration numbers already exist in the database
	existingRegNums, err := s.carRepo.GetExistingRegNums(db) // Fix here
	if err != nil {
		log.Println("[ERROR] Error retrieving existing registration numbers from the database:", err)
		return nil, err
	}

	// Debug log for existing registration numbers
	log.Println("[DEBUG] Existing registration numbers in the database:", existingRegNums)

	// Filter out registration numbers that already exist
	newRegNums := s.filterNewRegNums(regNums, existingRegNums)

	// If all registration numbers already exist, return an error
	if len(newRegNums) == 0 {
		log.Println("[INFO] All registration numbers already exist in the database")
		return nil, errors.New("all registration numbers already exist in the database")
	}

	// Info log for new registration numbers
	log.Println("[INFO] New registration numbers to be inserted into the database:", newRegNums)

	// Make a request to the external API with new registration numbers
	cars, err := s.carRepo.CreateCar(newRegNums, db) // Fix here
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

func (s *CarServiceImpl) GetCarsFiltered(criteria string, limit, offset int) ([]*model.Car, []*model.People, error) {
	// Call the repository method to get cars based on the provided criteria
	return s.carRepo.GetCarsFiltered(criteria, limit, offset) // Fix here
}

func (s *CarServiceImpl) UpdateCar(carID string, updatedCar model.Car) error {
	return s.carRepo.UpdateCar(carID, updatedCar) // Fix here
}

func (s *CarServiceImpl) DeleteCar(carID string) error {
	return s.carRepo.DeleteCar(carID) // Fix here
}

func (s *CarServiceImpl) filterNewRegNums(regNums, existingRegNums []string) []string {
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
			log.Print("[INFO] Registration number already exists in database:", regNum)
		}
	}
	return newRegNums
}
