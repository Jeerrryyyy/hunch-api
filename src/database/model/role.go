package model

import (
	"errors"
	"hunch-api/src/database"
)

type Role struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

func (role *Role) CreateRole() (*Role, error) {
	err := database.CONNECTION.Create(&role).Error

	if err != nil {
		return &Role{}, err
	}

	return role, nil
}

func GetRoleByName(name string) (Role, error) {
	var role Role

	if err := database.CONNECTION.Model(Role{}).Where("name = ?", name).Take(&role).Error; err != nil {
		return role, errors.New("role not found")
	}

	return role, nil
}
