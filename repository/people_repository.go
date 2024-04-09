package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

func (pr *PeopleRepository) CreatePerson(person model.People) (int, error) {
	query := "INSERT INTO people (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id"
	var personID int
	err := pr.db.QueryRow(query, person.Name, person.Surname, person.Patronymic).Scan(&personID)
	if err != nil {
		return 0, fmt.Errorf("error creating person: %v", err)
	}
	log.Printf("[INFO] Created person with ID %d", personID)
	return personID, nil
}

func (pr *PeopleRepository) GetPersonByID(personID int) (*model.People, error) {
	query := "SELECT name, surname, patronymic FROM people WHERE id = $1"
	row := pr.db.QueryRow(query, personID)
	var person model.People
	err := row.Scan(&person.Name, &person.Surname, &person.Patronymic)
	if err != nil {
		return nil, fmt.Errorf("error getting person by ID: %v", err)
	}
	log.Printf("[INFO] Retrieved person with ID %d", personID)
	return &person, nil
}

func (pr *PeopleRepository) UpdatePerson(personID int, updatedPerson model.People) error {
	setClause := " SET "
	var assignments []string

	if updatedPerson.Name != "" {
		assignments = append(assignments, fmt.Sprintf("name = '%s'", updatedPerson.Name))
	}
	if updatedPerson.Surname != "" {
		assignments = append(assignments, fmt.Sprintf("surname = '%s'", updatedPerson.Surname))
	}
	if updatedPerson.Patronymic != "" {
		assignments = append(assignments, fmt.Sprintf("patronymic = '%s'", updatedPerson.Patronymic))
	}

	setClause += strings.Join(assignments, ", ")

	query := "UPDATE people" + setClause + " WHERE id = $1"

	_, err := pr.db.Exec(query, personID)
	if err != nil {
		log.Printf("[ERROR] Failed to update person with ID %d: %v", personID, err)
		return err
	}
	log.Printf("[INFO] Updated person with ID %d", personID)
	return nil
}

func (pr *PeopleRepository) DeletePerson(personID int) error {
	query := "DELETE FROM people WHERE id = $1"
	_, err := pr.db.Exec(query, personID)
	if err != nil {
		return fmt.Errorf("error deleting person: %v", err)
	}
	log.Printf("[INFO] Deleted person with ID %d", personID)
	return nil
}
