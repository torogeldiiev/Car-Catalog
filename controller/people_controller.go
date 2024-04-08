package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/torogeldiiev/car_catalog/model"
	"github.com/torogeldiiev/car_catalog/service"
)

// PeopleController handles HTTP requests related to peop
type PeopleController struct {
	peopleService *service.PeopleService
}

// NewPeopleController creates a new instance of PeopleController
func NewPeopleController(peopleService *service.PeopleService) *PeopleController {
	return &PeopleController{
		peopleService: peopleService,
	}
}

// CreatePersonHandler handles requests to create a new person

// CreatePersonHandler @Summary Create person
// @Description Create a new person
// @Tags people
// @Accept json
// @Produce json
// @Param person body model.People true "Person object"
// @Success 201 {object} map[string]int "Created person ID"
// @Router /people/create [post]
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

// GetPersonByIDHandler @Summary Get person by ID
// @Description Get a person by their ID
// @Tags people
// @Accept json
// @Produce json
// @Param id query string true "Person ID"
// @Success 200 {object} model.People "Found person"
// @Router /people/get [get]
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

// UpdatePersonHandler handles requests to update an existing person

// UpdatePersonHandler @Summary Update person
// @Description Update an existing person
// @Tags people
// @Accept json
// @Produce json
// @Param id query string true "Person ID"
// @Param person body model.People true "Updated person object"
// @Success 200 "OK"
// @Router /people/update [put]
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

// DeletePersonHandler handles requests to delete a person by their ID

// DeletePersonHandler @Summary Delete person
// @Description Delete a person by their ID
// @Tags people
// @Accept json
// @Produce json
// @Param id query string true "Person ID"
// @Success 200 "OK"
// @Router /people/delete [delete]
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
