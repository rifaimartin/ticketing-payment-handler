package transaction

import (
	"ticketing-payment-handler/ticket"
	"ticketing-payment-handler/user"
	"time"

	"github.com/leekchan/accounting"
)

//Transaction struct
type Transaction struct {
	ID         int
	TicketID   int
	UserID     int
	Amount     int 
	Qty		   int
	Status     string
	Code       string
	User       user.User
	Ticket     ticket.Ticket
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//AmountFormatIDR fomat idr
func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "IDR ", Precision: 0, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}



