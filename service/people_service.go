package service

import (
	"fmt"
	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/repository"
	"log"
)

type PeopleService struct {
	peopleRepository *repository.PeopleRepository
}

// NewPeopleService creates a new instance of PeopleService
func NewPeopleService(peopleRepo *repository.PeopleRepository) *PeopleService {
	return &PeopleService{
		peopleRepository: peopleRepo,
	}
}

func (ps *PeopleService) CreatePerson(person model.People) (int, error) {
	log.Println("[INFO] Creating new person...")
	personID, err := ps.peopleRepository.CreatePerson(person)
	if err != nil {
		log.Printf("[ERROR] Error creating person: %v", err)
		return 0, fmt.Errorf("error creating person: %v", err)
	}
	log.Printf("[INFO] Successfully created person with ID: %d", personID)
	return personID, nil
}

func (ps *PeopleService) GetPersonByID(personID int) (*model.People, error) {
	log.Printf("[INFO] Retrieving person with ID: %d", personID)
	person, err := ps.peopleRepository.GetPersonByID(personID)
	if err != nil {
		log.Printf("[ERROR] Error retrieving person with ID %d: %v", personID, err)
		return nil, fmt.Errorf("error retrieving person with ID %d: %v", personID, err)
	}
	log.Printf("[INFO] Successfully retrieved person with ID %d", personID)
	return person, nil
}

func (ps *PeopleService) UpdatePerson(personID int, updatedPerson model.People) error {
	log.Printf("[INFO] Updating person with ID: %d", personID)
	err := ps.peopleRepository.UpdatePerson(personID, updatedPerson)
	if err != nil {
		log.Printf("[ERROR] Error updating person with ID %d: %v", personID, err)
		return fmt.Errorf("error updating person with ID %d: %v", personID, err)
	}
	log.Printf("[INFO] Successfully updated person with ID %d", personID)
	return nil
}

func (ps *PeopleService) DeletePerson(personID int) error {
	log.Printf("[INFO] Deleting person with ID: %d", personID)
	err := ps.peopleRepository.DeletePerson(personID)
	if err != nil {
		log.Printf("[ERROR] Error deleting person with ID %d: %v", personID, err)
		return fmt.Errorf("error deleting person with ID %d: %v", personID, err)
	}
	log.Printf("[INFO] Successfully deleted person with ID %d", personID)
	return nil
}
