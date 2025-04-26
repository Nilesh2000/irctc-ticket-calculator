package ticketpricecalculator

import (
	"slices"
)

// FixedTicketPriceCalculator is a struct that implements the TicketPriceCalculator interface.
// It calculates the ticket price based on a fixed price per stop for each coach type.
type FixedTicketPriceCalculator struct {
	train train
}

// NewFixedTicketPriceCalculator creates a new TicketPriceCalculator with the given route.
func NewFixedTicketPriceCalculator(number string, route []Station) *FixedTicketPriceCalculator {
	return &FixedTicketPriceCalculator{
		train: train{
			number: number,
			route:  route,
		},
	}
}

// CalculateTicketPrice calculates the ticket price based on the route.
func (s *FixedTicketPriceCalculator) Calculate(req *TicketBookingRequest) (totalPrice int, err error) {
	if err = s.validateRequest(req); err != nil {
		return
	}

	startIndex := slices.Index(s.train.route, req.StartStation)
	if startIndex == -1 {
		err = ErrInvalidStartStation
		return
	}
	endIndex := slices.Index(s.train.route, req.EndStation)
	if endIndex == -1 {
		err = ErrInvalidEndStation
		return
	}

	if startIndex >= endIndex {
		err = ErrInvalidStationOrder
		return
	}

	numStops := endIndex - startIndex

	var pricePerStop int
	switch req.Coach {
	case General:
		pricePerStop = GeneralCoachBasePrice
	case Sleeper:
		pricePerStop = SleeperCoachBasePrice
	default:
		err = ErrInvalidCoachType
		return
	}

	totalPrice = pricePerStop * numStops * req.NumPassengers
	return
}

func (s *FixedTicketPriceCalculator) validateRequest(req *TicketBookingRequest) error {
	if req.TrainNumber == "" {
		return ErrEmptyTrainNumber
	}
	if req.StartStation == "" {
		return ErrEmptyStartStation
	}
	if req.EndStation == "" {
		return ErrEmptyEndStation
	}
	if req.Coach == "" {
		return ErrEmptyCoachType
	}
	if req.NumPassengers < 1 {
		return ErrPassengersLessThanOne
	}
	return nil
}
