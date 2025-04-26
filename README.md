# Ticket Price Calculator

A Go module for calculating train ticket prices based on a fixed per-station pricing strategy. This module is designed to be used as an extension for various applications that need to calculate train ticket prices.

## Installation

```bash
go get github.com/Nilesh2000/ticketpricecalculator
```

## Features

- Calculate ticket prices based on the number of stops between stations
- Support for different coach types (General, Sleeper)
- Fixed price per stop for each coach type
- Comprehensive input validation
- Error handling for various edge cases

## Usage

```go
package main

import (
    "fmt"
    "github.com/Nilesh2000/ticketpricecalculator"
)

func main() {
    // Define the train route
    route := []ticketpricecalculator.Station{
        "StationA",
        "StationB",
        "StationC",
        "StationD",
    }

    // Create a new ticket price calculator with the train route
    calculator := ticketpricecalculator.NewFixedTicketPriceCalculator("12345", route)

    // Create a booking request
    request := &ticketpricecalculator.TicketBookingRequest{
        TrainNumber:   "12345",
        StartStation:  "StationA",
        EndStation:    "StationC",
        Coach:         ticketpricecalculator.Sleeper,
        NumPassengers: 2,
    }

    // Calculate the price
    price, err := calculator.Calculate(request)
    if err != nil {
        fmt.Printf("Error calculating price: %v\n", err)
        return
    }

    fmt.Printf("Total price: %d\n", price)
}
```

## API Reference

### Types

#### `TicketBookingRequest`

```go
type TicketBookingRequest struct {
    TrainNumber   string
    StartStation  Station
    EndStation    Station
    Coach         CoachType
    NumPassengers int
}
```

#### `CoachType`

```go
type CoachType string

const (
    General CoachType = "General"
    Sleeper CoachType = "Sleeper"
)
```

### Interface

#### `TicketPriceCalculator`

```go
type TicketPriceCalculator interface {
    Calculate(req *TicketBookingRequest) (int, error)
}
```

## Pricing Strategy

The ticket price is calculated based on:
- Number of stops between start and end stations
- Coach type (General or Sleeper)
- Number of passengers

Base prices per stop:
- General Coach: 20 units per stop
- Sleeper Coach: 40 units per stop

Total price = (Price per stop) × (Number of stops) × (Number of passengers)

## Error Handling

The calculator validates input and returns appropriate errors for:
- Empty train number
- Empty start/end station
- Empty coach type
- Invalid number of passengers (less than 1)
- Invalid start/end station (not in route)
- Invalid station order (start station after end station)
- Invalid coach type

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the LICENSE file in the root of this repository.
