package service

import (
	"database/sql"
	"github.com/torogeldiiev/car_catalog/model"
)

// CarService interface
type CarService interface {
	CreateCar(regNums []string, db *sql.DB) ([]*model.Car, error)
	UpdateCar(carID string, updatedCar model.Car) error
	DeleteCar(carID string) error
	GetCarsFiltered(criteria string, limit, offset int) ([]*model.Car, []*model.People, error)
}
