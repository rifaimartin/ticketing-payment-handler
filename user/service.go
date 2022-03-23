package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Service is interface
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

// NewService e
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Gender = input.Gender
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	user.Alamat = input.Alamat

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
