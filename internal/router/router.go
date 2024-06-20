package router

import (
	"github.com/TommyWoodley/flight-service/internal/api"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/search", api.SearchFlights()).Methods("POST")

	return r
}
