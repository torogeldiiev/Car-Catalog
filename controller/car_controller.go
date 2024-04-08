package controller

import "github.com/torogeldiiev/car_catalog/model"

// CarController interface
type CarController interface {
	CreateCar(regNums []string) ([]*model.Car, error)
	UpdateCar(carID string, updatedCar model.Car) error
	DeleteCar(carID string) error
	GetCarsFiltered(criteria string, limit, offset int) ([]*model.Car, error)
}
