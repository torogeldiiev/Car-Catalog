package repository

import (
	"database/sql"
	"fmt"

	"github.com/torogeldiiev/car_catalog/model"
)

// PeopleRepository handles interactions with the people table in the database
type PeopleRepository struct {
	db *sql.DB
}

// NewPeopleRepository creates a new instance of PeopleRepository
func NewPeopleRepository(db *sql.DB) *PeopleRepository {
	return &PeopleRepository{
		db: db,
	}
}

// CreatePerson inserts a new person into the database
func (pr *PeopleRepository) CreatePerson(person model.People) (int, error) {
	query := "INSERT INTO people (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id"
	var personID int
	err := pr.db.QueryRow(query, person.Name, person.Surname, person.Patronymic).Scan(&personID)
	if err != nil {
		return 0, fmt.Errorf("error creating person: %v", err)
	}
	return personID, nil
}

// GetPersonByID retrieves a person from the database by their ID
func (pr *PeopleRepository) GetPersonByID(personID int) (*model.People, error) {
	query := "SELECT name, surname, patronymic FROM people WHERE id = $1"
	row := pr.db.QueryRow(query, personID)
	var person model.People
	err := row.Scan(&person.Name, &person.Surname, &person.Patronymic)
	if err != nil {
		return nil, fmt.Errorf("error getting person by ID: %v", err)
	}
	return &person, nil
}

// UpdatePerson updates an existing person in the database
func (pr *PeopleRepository) UpdatePerson(personID int, updatedPerson model.People) error {
	query := "UPDATE people SET name = $1, surname = $2, patronymic = $3 WHERE id = $4"
	_, err := pr.db.Exec(query, updatedPerson.Name, updatedPerson.Surname, updatedPerson.Patronymic, personID)
	if err != nil {
		return fmt.Errorf("error updating person: %v", err)
	}
	return nil
}

// DeletePerson deletes a person from the database by their ID
func (pr *PeopleRepository) DeletePerson(personID int) error {
	query := "DELETE FROM people WHERE id = $1"
	_, err := pr.db.Exec(query, personID)
	if err != nil {
		return fmt.Errorf("error deleting person: %v", err)
	}
	return nil
}
