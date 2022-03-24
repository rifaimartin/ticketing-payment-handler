package ticket

//Service interface
type Service interface {
	GetTickets(userID int) ([]Ticket, error)
	CreateTicket(input CreateTicketInput) (Ticket, error)
}

type service struct {
	repository Repository
}

//NewService initiaze newservice
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTickets(userID int) ([]Ticket, error) {
	// if userID != 0 {
	// 	tickets, err := s.repository.FindByUserID(userID)
	// 	if err != nil {
	// 		return tickets, err
	// 	}

	// 	return tickets, nil
	// }

	tickets, err := s.repository.FindAll()
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func (s *service) CreateTicket(input CreateTicketInput) (Ticket, error) {
	ticket := Ticket{}
	ticket.Name = input.Name
	ticket.ShortDescription = input.ShortDescription
	ticket.Qty = input.Qty
	ticket.price = input.Price

	newTicket, err := s.repository.Save(ticket)
	if err != nil {
		return newTicket, err
	}

	return newTicket, nil
}