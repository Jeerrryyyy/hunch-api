package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(userId uint) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_MINUTE_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["tokenType"] = "access"
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(tokenLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func GenerateRefreshToken(userId uint) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["tokenType"] = "refresh"
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ValidateToken(context *gin.Context) error {
	tokenString := ExtractToken(context)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(context *gin.Context) string {
	token := context.Query("token")

	if token != "" {
		return token
	}

	bearerToken := context.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(context *gin.Context) (uint, error) {
	tokenString := ExtractToken(context)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)

		if err != nil {
			return 0, err
		}

		return uint(id), nil
	}

	return 0, nil
}

func ExtractTokenType(context *gin.Context) (string, error) {
	tokenString := ExtractToken(context)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return "unknown", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		tokenType := fmt.Sprintf("%s", claims["tokenType"])

		if err != nil {
			return "unknown", err
		}

		return tokenType, nil
	}

	return "unknown", nil
}
