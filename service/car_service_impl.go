package service

import (
	"database/sql"
	"errors"
	"fmt"
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
	existingRegNums, err := s.carRepo.GetExistingRegNums(db)
	if err != nil {
		log.Println("[ERROR] Error retrieving existing registration numbers from the database:", err)
		return nil, err
	}

	log.Println("[DEBUG] Existing registration numbers in the database:", existingRegNums)

	// Filter out registration numbers that already exist
	newRegNums := s.filterNewRegNums(regNums, existingRegNums)

	if len(newRegNums) == 0 {
		log.Println("[INFO] All registration numbers already exist in the database")
		return nil, errors.New("all registration numbers already exist in the database")
	}

	log.Println("[INFO] New registration numbers to be inserted into the database:", newRegNums)

	// Make a request to the external API with new registration numbers
	cars, err := s.carRepo.CreateCar(newRegNums, db)
	if err != nil {
		log.Println("[ERROR] Error creating cars in the database:", err)
		return nil, err
	}

	for _, car := range cars {
		log.Println("[INFO] Successfully inserted car with registration number:", car.RegNum)
	}

	log.Println("[DEBUG] Exiting CreateCar service method")

	return cars, nil
}

func (s *CarServiceImpl) GetCarsFiltered(criteria string, limit, offset int) ([]*model.Car, []*model.People, error) {
	return s.carRepo.GetCarsFiltered(criteria, limit, offset)
}

func (s *CarServiceImpl) UpdateCar(carID string, updatedCar model.Car) error {
	return s.carRepo.UpdateCar(carID, updatedCar)
}

func (s *CarServiceImpl) DeleteCar(carID string) error {
	err := s.carRepo.DeleteCar(carID)
	if err != nil {
		log.Printf("[ERROR] Error deleting car with ID %s: %v", carID, err)
		return fmt.Errorf("error deleting car with ID %s: %v", carID, err)
	}
	log.Printf("[INFO] Successfully deleted car with ID %s", carID)
	return nil
}

func (s *CarServiceImpl) filterNewRegNums(regNums, existingRegNums []string) []string {
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
