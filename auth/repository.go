package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	Create(song User) (User, error)
	SignIn(signin SignIn) (User, error)
	UpdateAddress(user User) (User, error)
	FindByEmail(email any) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return user, err
}

func (r *repository) FindAll() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return user, err
}

func (r *repository) SignIn(signin SignIn) (User, error) {
	var user User
	err := r.db.Debug().Where(&User{Email: signin.Email}).First(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR LOGIN1")
		println("=====================")
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signin.Password))

	if err != nil {
		println("=====================")
		println("ERROR LOGIN2")
		println("=====================")
	}

	return user, err
}

func (r *repository) UpdateAddress(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}
	return user, err
}

func (r *repository) FindByEmail(email any) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE FB-EMAIL")
		println("=====================")
	}

	return user, err
}
