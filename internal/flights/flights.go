package flights

import (
	"github.com/TommyWoodley/flight-service/internal/models"
)

// FetchFlightDataFromAPI returns dummy flight data for testing purposes
func FetchFlightDataFromAPI(departureTime, returnTime string) ([]models.Flight, error) {
	// Dummy data for demonstration purposes
	flights := []models.Flight{
		{
			Flight:        "Flight 1",
			DepartureTime: departureTime,
			ReturnTime:    returnTime,
			Origin:        "Heathrow",
			Destination:   "Paris",
			Price:         "$200",
			Duration:      "2h 30m",
		},
		{
			Flight:        "Flight 2",
			DepartureTime: departureTime,
			ReturnTime:    returnTime,
			Origin:        "Heathrow",
			Destination:   "Berlin",
			Price:         "$250",
			Duration:      "3h",
		},
	}
	return flights, nil
}
