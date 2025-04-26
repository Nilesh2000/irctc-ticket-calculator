package ticketpricecalculator

import (
	"errors"
	"testing"
)

// TestTicketPriceCalculator_CalculateTicketPrice tests the CalculateTicketPrice method of TicketPriceCalculator.
func TestTicketPriceCalculator_CalculateTicketPrice(t *testing.T) {
	tests := []struct {
		name          string
		request       TicketBookingRequest
		expectedPrice int
		expectedError error
	}{
		{
			name:          "Karjat to Chinchwad - General - 1 passenger",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Karjat", EndStation: "Chinchwad", Coach: General, NumPassengers: 1},
			expectedPrice: 40,
			expectedError: nil,
		},
		{
			name:          "Mumbai to Pune - Sleeper - 3 passengers",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Mumbai", EndStation: "Pune", Coach: Sleeper, NumPassengers: 3},
			expectedPrice: 480,
			expectedError: nil,
		},
		{
			name:          "Number of passengers is zero",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Karjat", EndStation: "Chinchwad", Coach: General, NumPassengers: 0},
			expectedPrice: 0,
			expectedError: ErrPassengersLessThanOne,
		},
		{
			name:          "Invalid start station",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "InvalidStation", EndStation: "Chinchwad", Coach: General, NumPassengers: 1},
			expectedPrice: 0,
			expectedError: ErrInvalidStartStation,
		},
		{
			name:          "Invalid end station",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Karjat", EndStation: "InvalidEndStation", Coach: General, NumPassengers: 1},
			expectedPrice: 0,
			expectedError: ErrInvalidEndStation,
		},
		{
			name:          "End station before start station",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Chinchwad", EndStation: "Karjat", Coach: General, NumPassengers: 1},
			expectedPrice: 0,
			expectedError: ErrInvalidStationOrder,
		},
		{
			name:          "Invalid coach type",
			request:       TicketBookingRequest{TrainNumber: "12345", StartStation: "Karjat", EndStation: "Chinchwad", Coach: "InvalidCoach", NumPassengers: 1},
			expectedPrice: 0,
			expectedError: ErrInvalidCoachType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator := NewFixedTicketPriceCalculator("12345", []Station{"Mumbai", "Karjat", "Lonavala", "Chinchwad", "Pune"})

			price, err := calculator.Calculate(&tt.request)

			if price != tt.expectedPrice {
				t.Errorf("expected price %d, got %d", tt.expectedPrice, price)
			}

			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
