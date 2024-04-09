package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/service"
)

type PeopleController struct {
	peopleService *service.PeopleService
}

// NewPeopleController creates a new instance of PeopleController
func NewPeopleController(peopleService *service.PeopleService) *PeopleController {
	return &PeopleController{
		peopleService: peopleService,
	}
}

func (pc *PeopleController) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var person model.People
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	personID, err := pc.peopleService.CreatePerson(person)
	if err != nil {
		http.Error(w, "Failed to create person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"personID": personID})
}

func (pc *PeopleController) GetPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	personIDStr := r.URL.Query().Get("id")
	if personIDStr == "" {
		http.Error(w, "Person ID is required", http.StatusBadRequest)
		return
	}

	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	person, err := pc.peopleService.GetPersonByID(personID)
	if err != nil {
		http.Error(w, "Failed to retrieve person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (pc *PeopleController) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	personIDStr := r.URL.Query().Get("id")
	if personIDStr == "" {
		http.Error(w, "Person ID is required", http.StatusBadRequest)
		return
	}

	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	var updatedPerson model.People
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := pc.peopleService.UpdatePerson(personID, updatedPerson); err != nil {
		http.Error(w, "Failed to update person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pc *PeopleController) DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	personIDStr := r.URL.Query().Get("id")
	if personIDStr == "" {
		http.Error(w, "Person ID is required", http.StatusBadRequest)
		return
	}

	personID, err := strconv.Atoi(personIDStr)
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	if err := pc.peopleService.DeletePerson(personID); err != nil {
		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
