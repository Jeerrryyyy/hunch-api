package controller

import (
	"github.com/gin-gonic/gin"
	"hunch-api/src/database/model"
	"hunch-api/src/util/token"
	"net/http"
)

func CurrentUser(context *gin.Context) {
	userId, err := token.ExtractTokenID(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

type RegisterInput struct {
	Email     string   `json:"email" binding:"required"`
	FirstName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	Roles     []string `json:"roles" binding:"required"`
}

func CreateUser(context *gin.Context) {
	var registerInput RegisterInput

	if err := context.ShouldBindJSON(&registerInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles := make([]model.Role, len(registerInput.Roles))
	for i, roleName := range registerInput.Roles {
		role, err := model.GetRoleByName(roleName)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roles[i] = role
	}

	user := model.User{
		Email:          registerInput.Email,
		FirstName:      registerInput.FirstName,
		LastName:       registerInput.LastName,
		Password:       registerInput.Password,
		ProfilePicture: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
		Roles:          roles,
	}

	_, err := user.CreateUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registered"})
}
