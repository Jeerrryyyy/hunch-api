package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"hunch-api/src/database"
	"strings"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	FirstName string `gorm:"size:255;not null;" json:"firstName"`
	LastName  string `gorm:"size:255;not null;" json:"lastName"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	Roles     []Role `gorm:"many2many:user_roles;" json:"roles"`
}

func (user *User) PrepareGive() {
	user.Password = "removed for security reasons"
}

func (user *User) CreateUser() (*User, error) {
	err := database.CONNECTION.Create(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) UpdateUser() (*User, error) {
	err := database.CONNECTION.Save(&user).Error

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) BeforeSave(_ *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}

func GetUserById(userId uint) (User, error) {
	var user User

	if err := database.CONNECTION.Model(User{}).Preload("Roles").First(&user, userId).Error; err != nil {
		return user, errors.New("user not found")
	}

	user.PrepareGive()

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User

	if err := database.CONNECTION.Model(User{}).Preload("Roles").Where("email = ?", email).Take(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	user.PrepareGive()

	return user, nil
}
