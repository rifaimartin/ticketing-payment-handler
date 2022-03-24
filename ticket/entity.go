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
	imageUrl 		 string
	Qty 		 	 int
	price            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// priceFormatIDR format price
func (t Ticket) priceFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.price)
}


