package ticketpricecalculator

import "errors"

var (
	// Validation errors
	ErrEmptyTrainNumber      = errors.New("train number cannot be empty")
	ErrEmptyStartStation     = errors.New("start station cannot be empty")
	ErrEmptyEndStation       = errors.New("end station cannot be empty")
	ErrEmptyCoachType        = errors.New("coach type cannot be empty")
	ErrPassengersLessThanOne = errors.New("number of passengers must be greater than zero")

	// Business logic errors
	ErrInvalidStartStation = errors.New("invalid start station")
	ErrInvalidEndStation   = errors.New("invalid end station")
	ErrInvalidStationOrder = errors.New("start station should come before end station")
	ErrInvalidCoachType    = errors.New("invalid coach type")
)
