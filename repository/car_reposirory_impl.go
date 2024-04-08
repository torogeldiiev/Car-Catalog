package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/torogeldiiev/car_catalog/database"
	"github.com/torogeldiiev/car_catalog/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// CarRepositoryImpl implements CarRepository interface
type CarRepositoryImpl struct {
	DB *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepositoryImpl {
	return &CarRepositoryImpl{DB: db}
}

func (r *CarRepositoryImpl) CreateCar(regNums []string, db *sql.DB) ([]*model.Car, error) {
	// Construct the URL for the external API request
	url := fmt.Sprintf("http://localhost:8081/info?regNums=%s", strings.Join(regNums, ","))

	// Make a GET request to the external API
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[INFO] Error making GET request to external API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		// Log the error message if the status code is not OK
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[INFO] Error response from external API: %s", string(body))
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response body into a slice of Car objects
	var cars []*model.Car
	if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
		log.Printf("[INFO] Error decoding response body: %v", err)
		return nil, err
	}

	// Insert the fetched cars into the database
	for _, car := range cars {
		// Prepare the SQL statement for inserting a new car
		stmt, err := db.Prepare("INSERT INTO cars (reg_num, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)")
		if err != nil {
			log.Printf("[INFO] Error preparing SQL statement: %v", err)
			return nil, err
		}
		defer stmt.Close()

		// Execute the prepared SQL statement with car details
		_, err = stmt.Exec(car.RegNum, car.Make, car.Model, car.Year, car.OwnerID)
		if err != nil {
			log.Println("[INFO] Error inserting car into database: %v", err)
			// Continue to the next car if an error occurs
			continue
		}
	}

	log.Printf("[INFO] Successfully inserted %d cars into the database", len(cars))

	return cars, nil
}

func (r *CarRepositoryImpl) UpdateCar(carID string, updatedCar model.Car) error {
	setClause := " SET "
	var assignments []string

	if updatedCar.Make != nil {
		assignments = append(assignments, fmt.Sprintf("mark = '%s'", *updatedCar.Make))
	}
	if updatedCar.Model != nil {
		assignments = append(assignments, fmt.Sprintf("model = '%s'", *updatedCar.Model))
	}
	if updatedCar.Year != nil {
		assignments = append(assignments, fmt.Sprintf("year = %d", *updatedCar.Year))
	}

	setClause += strings.Join(assignments, ", ")

	query := "UPDATE cars" + setClause + " WHERE id = $1"

	_, err := database.DB.Exec(query, carID)
	if err != nil {
		log.Printf("[INFO] Error updating car: %v", err) // Log the error
		return err
	}
	return nil
}
func (r *CarRepositoryImpl) DeleteCar(carID string) error {
	query := "DELETE FROM cars WHERE id = $1"
	_, err := database.DB.Exec(query, carID)
	if err != nil {
		log.Printf("[INFO] Error deleting car: %v", err) // Log the error
		return err
	}

	return nil
}

func (r *CarRepositoryImpl) GetExistingRegNums(db *sql.DB) ([]string, error) {
	// Query the database to get existing registration numbers
	rows, err := db.Query("SELECT reg_num FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var existingRegNums []string
	// Iterate through the rows and collect registration numbers
	for rows.Next() {
		var regNum string
		if err := rows.Scan(&regNum); err != nil {
			return nil, err
		}
		existingRegNums = append(existingRegNums, regNum)
	}
	return existingRegNums, nil
}

func (r *CarRepositoryImpl) GetCarsFiltered(criteria string, limit, offset int) ([]*model.Car, []*model.People, error) {
	queryBuilder := squirrel.Select("cars.id", "cars.reg_num", "cars.mark", "cars.model", "cars.year", "cars.owner_id", "people.name", "people.surname", "people.patronymic").
		From("cars").
		Join("people ON cars.owner_id = people.id").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	// Adding the WHERE clause based on the provided criteria
	if criteria != "" {
		queryBuilder = queryBuilder.Where(criteria)
	}

	// Generate the final SQL query
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Println("[INFO] Error generating SQL query: %v", err)
		return nil, nil, err
	}

	// Execute the SQL query
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		log.Printf("[INFO] Error executing SQL query: %v", err)
		return nil, nil, err
	}
	defer rows.Close()

	// Iterate through the rows and collect cars and their owners
	var cars []*model.Car
	var owners []*model.People
	for rows.Next() {
		var car model.Car
		var owner model.People
		err := rows.Scan(&car.ID, &car.RegNum, &car.Make, &car.Model, &car.Year, &car.OwnerID, &owner.Name, &owner.Surname, &owner.Patronymic)
		if err != nil {
			log.Printf("[INFO] Error scanning row: %v", err)
			return nil, nil, err
		}
		cars = append(cars, &car)
		owners = append(owners, &owner)
	}

	return cars, owners, nil
}
