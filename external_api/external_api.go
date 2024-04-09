package external_api

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"

	"github.com/torogeldiiev/car_catalog/model"
)

// MockServer represents a mock server for the external API
type MockServer struct{}

// NewMockServer creates a new instance of the mock server
func NewMockServer() *MockServer {
	return &MockServer{}
}

// Start starts the mock server
func (s *MockServer) Start() {
	http.HandleFunc("/info", s.infoHandler)
	http.ListenAndServe(":8081", nil)
}

// infoHandler handles the /info endpoint
// затычка реализующая роль внешнего апи
func (s *MockServer) infoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the regNums parameter from the query string
	regNums := r.URL.Query()["regNums"]

	log.Printf("Processing request with registration numbers: %v", regNums)

	var cars []*model.Car
	for _, regNum := range regNums {
		car := model.Car{
			RegNum:  &regNum,
			Make:    stringPointer("Lada"),
			Model:   stringPointer("Vesta"),
			Year:    intPointer(2002),
			OwnerID: intPointer(1), // Convert int constant to a pointer to int
		}
		cars = append(cars, &car)
	}

	responseJSON, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response status code to 200 (OK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}
