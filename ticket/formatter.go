package ticket

//TicketFormatter type formatter
type TicketFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Qty              int `json:"qty"`
	Price            int    `json:"price"`
}

//FormatTicket this function format ticket return
func FormatTicket(ticket Ticket) TicketFormatter {
	ticketFormatter := TicketFormatter{}
	ticketFormatter.ID = ticket.ID
	ticketFormatter.Name = ticket.Name
	ticketFormatter.ShortDescription = ticket.ShortDescription
	ticketFormatter.Qty = ticket.Qty
	ticketFormatter.Price = ticket.price

	return ticketFormatter
}

// FormatTicket format data tickets
func FormatTickets(tickets []Ticket) []TicketFormatter {

	ticketsFormatter := []TicketFormatter{}
	//
	for _, ticket := range tickets {
		ticketFormatter := FormatTicket(ticket)
		ticketsFormatter = append(ticketsFormatter, ticketFormatter)
	}

	return ticketsFormatter
}
