package ticketpricecalculator

// Station represents a train station.
type Station string

// CoachType represents the type of coach (General, Sleeper, etc.) in a train.
type CoachType string

const (
	General CoachType = "General"
	Sleeper CoachType = "Sleeper"
)

const (
	GeneralCoachBasePrice int = 20
	SleeperCoachBasePrice int = 40
)

// train represents a train with a number and an ordered route.
type train struct {
	number string
	// route specifies the list of the stations in the order they are visited.
	route []Station
}

// TicketBookingRequest represents a request to book a ticket.
type TicketBookingRequest struct {
	TrainNumber   string
	StartStation  Station
	EndStation    Station
	Coach         CoachType
	NumPassengers int
}
