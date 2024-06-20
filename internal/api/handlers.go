package api

import (
	"encoding/json"
	"github.com/TommyWoodley/flight-service/internal/flights"
	"github.com/TommyWoodley/flight-service/internal/models"
	"net/http"
)

func SearchFlights() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode Client Request into FlightSearchRequest
		var request models.FlightSearchRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		availableFlights, err := flights.FetchFlightDataFromAPI(request.DepartureTime, request.ReturnTime)
		if err != nil {
			http.Error(w, "Error fetching flight data", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(availableFlights)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}
}
