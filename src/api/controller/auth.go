package controller

import (
	"github.com/gin-gonic/gin"
	"hunch-api/src/api/service"
	"hunch-api/src/util/token"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(context *gin.Context) {
	var loginInput LoginInput

	if err := context.ShouldBindJSON(&loginInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := service.LoginCheck(loginInput.Email, loginInput.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is incorrect"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func ValidateToken(context *gin.Context) {
	err := token.ValidateToken(context)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{"valid": false, "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"valid": true, "error": nil})
}

func RefreshToken(context *gin.Context) {

}
