package transaction

import (
	"fmt"
	"strconv"
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
	ProcessPayment(input TransactionNotificationInput) error
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

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)
	
	fmt.Println("i am here")
	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	fmt.Println("you here")
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	ticket, err := s.ticketRepository.FindByID(updatedTransaction.TicketID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		ticket.Qty = ticket.Qty - 1

		_, err := s.ticketRepository.Update(ticket)
		if err != nil {
			return err
		}

	}

	return nil

}
