package transaction

import (
	"ticketing-payment-handler/ticket"
)

type service struct {
	repository         Repository
	ticketRepository   ticket.Repository
}

//Service interface
type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

//NewService newservice
func NewService(repository Repository, ticketRepository ticket.Repository) *service {
	return &service{repository, ticketRepository}
}


func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.TicketID = input.TicketID
	transaction.Amount = input.Amount
	transaction.Qty = input.Qty
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
