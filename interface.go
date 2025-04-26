package ticketpricecalculator

// TicketPriceCalculator is an interface that defines the method to calculate ticket prices.
type TicketPriceCalculator interface {
	Calculate(req *TicketBookingRequest) (int, error)
}
