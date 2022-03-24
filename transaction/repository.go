package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

//Repository interface
type Repository interface {
	GetByID(ID int) (Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

//NewRepository semothing
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}


func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction

	// strict ilmu
	err := r.db.Preload("Ticket").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
