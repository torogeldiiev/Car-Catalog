package main

import (
	"github.com/torogeldiiev/car_catalog/controller"
	"github.com/torogeldiiev/car_catalog/database"
	"github.com/torogeldiiev/car_catalog/external_api"
	"github.com/torogeldiiev/car_catalog/repository"
	"github.com/torogeldiiev/car_catalog/service"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	// Initialize repository
	carRepo := repository.NewCarRepository(db)
	peopleRepo := repository.NewPeopleRepository(db)

	// Initialize services
	carService := service.NewCarService(carRepo)
	peopleService := service.NewPeopleService(peopleRepo)

	// Initialize controllers
	carController := controller.NewCarController(carService, db)
	peopleController := controller.NewPeopleController(peopleService)

	// Register routes
	http.HandleFunc("/cars/create", carController.CreateCarHandler)
	http.HandleFunc("/cars/get", carController.GetCarByIDHandler)
	http.HandleFunc("/cars/update", carController.UpdateCarHandler)
	http.HandleFunc("/cars/delete", carController.DeleteCarHandler)

	http.HandleFunc("/people/create", peopleController.CreatePersonHandler)
	http.HandleFunc("/people/get", peopleController.GetPersonByIDHandler)
	http.HandleFunc("/people/update", peopleController.UpdatePersonHandler)
	http.HandleFunc("/people/delete", peopleController.DeletePersonHandler)

	// Create and start the mock server
	mockServer := external_api.NewMockServer()
	go mockServer.Start()

	// Start the main server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
