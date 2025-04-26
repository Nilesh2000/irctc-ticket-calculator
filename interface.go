package ticketpricecalculator

// TicketPriceCalculator is an interface that defines the method to calculate ticket prices.
type TicketPriceCalculator interface {
	// Calculate calculates the ticket price for a given request.
	Calculate(req *TicketBookingRequest) (int, error)
}
