package controller

import (
	"github.com/gin-gonic/gin"
	"hunch-api/src/api/service"
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
		context.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func RefreshToken(context *gin.Context) {

}
