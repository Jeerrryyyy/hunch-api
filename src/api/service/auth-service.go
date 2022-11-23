package service

import (
	"golang.org/x/crypto/bcrypt"
	"hunch-api/src/database"
	"hunch-api/src/database/model"
	"hunch-api/src/util/token"
)

func LoginCheck(email string, password string) (string, string, error) {
	user := model.User{}

	err := database.CONNECTION.Model(model.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", err
	}

	accessToken, refreshToken, err := GenerateTokens(user.ID)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateTokens(userId uint) (string, string, error) {
	accessToken, err := token.GenerateToken(userId)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := token.GenerateRefreshToken(userId)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
