package ticket

import (
	"gorm.io/gorm"
)

//Repository interface
type Repository interface {
	FindAll() ([]Ticket, error)
	Save(ticket Ticket) (Ticket, error)
	FindByID(ID int) (Ticket, error)
	Update(ticket Ticket) (Ticket, error)
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

func (r *repository) FindByID(ID int) (Ticket, error) {
	var ticket Ticket
	err := r.db.Where("id = ?", ID).Find(&ticket).Error

	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (r *repository) Update(ticket Ticket) (Ticket, error) {
	err := r.db.Save(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}
