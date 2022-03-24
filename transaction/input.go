package transaction

import "ticketing-payment-handler/user"


type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	TicketID   int `json:"ticket_id" binding:"required"`
	Qty        int `json:"qty" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
