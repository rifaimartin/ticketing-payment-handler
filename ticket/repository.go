package ticket

import (
	"gorm.io/gorm"
)

//Repository interface
type Repository interface {
	FindAll() ([]Ticket, error)
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

