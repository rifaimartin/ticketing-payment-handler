package ticket

//Service interface
type Service interface {
	GetTickets(userID int) ([]Ticket, error)
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
