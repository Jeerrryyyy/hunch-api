package model

import (
	"errors"
	"hunch-api/src/database"
)

type RefreshToken struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Token       string `gorm:"size:255;not null;unique" json:"token"`
	Invalidated bool   `gorm:"not null" json:"invalidated"`
}

func (refreshToken *RefreshToken) CreateToken() (*RefreshToken, error) {
	err := database.CONNECTION.Create(&refreshToken).Error

	if err != nil {
		return &RefreshToken{}, err
	}

	return refreshToken, nil
}

func (refreshToken *RefreshToken) UpdateToken() (*RefreshToken, error) {
	err := database.CONNECTION.Save(&refreshToken).Error

	if err != nil {
		return &RefreshToken{}, err
	}

	return refreshToken, nil
}

func GetRefreshTokenByToken(queryToken string) (RefreshToken, error) {
	var refreshToken RefreshToken

	if err := database.CONNECTION.Model(RefreshToken{}).Where("token = ?", queryToken).Take(&refreshToken).Error; err != nil {
		return refreshToken, errors.New("refreshToken not found")
	}

	return refreshToken, nil
}
