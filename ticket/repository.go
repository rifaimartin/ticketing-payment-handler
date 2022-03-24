package ticket

import (
	"gorm.io/gorm"
)

//Repository interface
type Repository interface {
	FindAll() ([]Ticket, error)
	Save(ticket Ticket) (Ticket, error)

}
type repository struct {
	db *gorm.DB
}

//NewRepository initiaze
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Ticket, error) {
	var tickets []Ticket
	err := r.db.Find(&tickets).Error
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func (r *repository) Save(ticket Ticket) (Ticket, error) {
	err := r.db.Create(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}
