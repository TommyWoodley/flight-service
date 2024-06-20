package models

type FlightSearchRequest struct {
	DepartureTime string `json:"departureTime"`
	ReturnTime    string `json:"returnTime"`
}

type Flight struct {
	Flight        string `json:"flight"`
	DepartureTime string `json:"departureTime"`
	ReturnTime    string `json:"returnTime"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Price         string `json:"price"`
	Duration      string `json:"duration"`
}
