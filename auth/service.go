package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]User, error)
	FindByEmail(email any) (User, error)
	Create(userInput UserInput) (User, error)
	SignIn(signin SignIn) (User, error)
	UpdateAddress(addressInput AddressInput, user_email string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(userInput UserInput) (User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	user := User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hash),
		Phone:    userInput.Phone,
		Type:     "buyer",
	}
	newuser, err := s.repository.Create(user)
	return newuser, err
}

func (s *service) SignIn(signin SignIn) (User, error) {

	user, err := s.repository.SignIn(signin)
	return user, err
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err

}

func (s *service) UpdateAddress(addressInput AddressInput, user_email string) (User, error) {

	user, err := s.repository.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
	}
	user.City = addressInput.City
	user.Address = addressInput.Address

	newaddress, err := s.repository.UpdateAddress(user)
	return newaddress, err
}

func (s *service) FindByEmail(email any) (User, error) {
	user, err := s.repository.FindByEmail(email)
	return user, err
}
