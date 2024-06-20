# Flight Search Backend

This project is a backend service for searching flights based on user-provided departure and return times. It fetches flight data from an external API (currently returns dummy data) and caches the results using Redis for optimized performance.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- Redis server (version 5.0 or higher)

### Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/your-username/flight-search-backend.git
   cd flight-search-backend

2. **Install dependencies**:
  ```bash
  go mod tidy
  ```

### Running the server

1. **Start the Go server**:
    ```bash
    go run cmd/server/main.go
    ```
2. The server will start on http://localhost:8000.

## API Endpoints
**Search Flights**
- Endpoint: /api/search
- Method: POST
- Description: Searches for flights based on the user's specified departure and return times. Caches the results for future requests.
- Request Payload:
  ```json
  {
    "departureTime": "2023-07-01T18:00:00Z",
    "returnTime": "2023-07-03T23:59:59Z"
  }
```
- Response Payload:
  ```json
  [
    {
      "flight": "Flight 1",
      "departureTime": "2023-07-01T18:00:00Z",
      "returnTime": "2023-07-03T23:59:59Z",
      "origin": "Heathrow",
      "destination": "Paris",
      "price": "$200",
      "duration": "2h 30m"
    },
    {
      "flight": "Flight 2",
      "departureTime": "2023-07-01T18:00:00Z",
      "returnTime": "2023-07-03T23:59:59Z",
      "origin": "Heathrow",
      "destination": "Berlin",
      "price": "$250",
      "duration": "3h"
    }
  ]
  ```

