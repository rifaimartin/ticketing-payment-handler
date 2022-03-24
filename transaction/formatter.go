package transaction

import (
	"time"
)

//TransactionFormatter struct
type TransactionFormatter struct {
	ID         int    `json:"id"`
	TicketID   int       `json:"ticket_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
}

// FormatTransaction single object
func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.TicketID = transaction.TicketID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	return formatter
}

//TicketTransactionFormatter struct
type TicketTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price    int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}


//UserTransactionFormatter struct
type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	Ticket    TicketFormatter   `json:"ticket"`
	CreatedAt time.Time         `json:"created_at"`
}

//TicketFormatter struct
type TicketFormatter struct {
	Name     string `json:"name"`
}

//FormatUserTransaction formatter function
func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	ticketFormatter := TicketFormatter{}
	ticketFormatter.Name = transaction.Ticket.Name

	formatter.Ticket = ticketFormatter

	return formatter
}

//FormatUserTransactions many object
func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
