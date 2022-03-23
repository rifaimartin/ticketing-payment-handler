package ticket

import (
	"time"

	"github.com/leekchan/accounting"
)

// ticket Struct
type Ticket struct {
	ID               int
	Name             string
	ShortDescription string
	qty 		 	 int
	price            int
	feeFinal		 int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// priceFormatIDR format goal amount
func (t Ticket) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.price)
}


