package service

import (
	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/repository"
)

// PeopleService handles business logic related to people entities
type PeopleService struct {
	peopleRepository *repository.PeopleRepository
}

// NewPeopleService creates a new instance of PeopleService
func NewPeopleService(peopleRepo *repository.PeopleRepository) *PeopleService {
	return &PeopleService{
		peopleRepository: peopleRepo,
	}
}

// CreatePerson creates a new person
func (ps *PeopleService) CreatePerson(person model.People) (int, error) {
	return ps.peopleRepository.CreatePerson(person)
}

// GetPersonByID retrieves a person by their ID
func (ps *PeopleService) GetPersonByID(personID int) (*model.People, error) {
	return ps.peopleRepository.GetPersonByID(personID)
}

// UpdatePerson updates an existing person
func (ps *PeopleService) UpdatePerson(personID int, updatedPerson model.People) error {
	return ps.peopleRepository.UpdatePerson(personID, updatedPerson)
}

// DeletePerson deletes a person by their ID
func (ps *PeopleService) DeletePerson(personID int) error {
	return ps.peopleRepository.DeletePerson(personID)
}
