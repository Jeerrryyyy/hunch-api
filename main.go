package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hunch-api/src/api/controller"
	"hunch-api/src/api/middleware"
	"hunch-api/src/database"
	"hunch-api/src/database/model"
	"log"
	"net/http"
)

func main() {
	database.ConnectDatabase()
	database.MigrateTables(&model.Role{}, &model.User{}, &model.RefreshToken{})

	createDefaults()

	err := setupRouter().Run()

	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	v1 := router.Group("/v1")

	v1.GET("/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	auth := v1.Group("/auth")
	auth.POST("/login", controller.LoginUser)

	user := v1.Group("/user")
	user.GET("/current", middleware.JwtAuthMiddleWare([]string{}), controller.CurrentUser)
	user.POST("/create", middleware.JwtAuthMiddleWare([]string{"ADMINISTRATOR"}), controller.CreateUser)

	return router
}

func createDefaults() {
	var count int64

	err := database.CONNECTION.Model(&model.Role{}).Count(&count).Error

	if err != nil {
		panic(err)
	}
	if count != 0 {
		return
	}

	roleAdministrator := model.Role{Name: "ADMINISTRATOR"}
	_, err = roleAdministrator.CreateRole()
	if err != nil {
		panic(err)
	}

	roleDeveloper := model.Role{Name: "DEVELOPER"}
	_, err = roleDeveloper.CreateRole()
	if err != nil {
		panic(err)
	}

	roleUser := model.Role{Name: "USER"}
	_, err = roleUser.CreateRole()
	if err != nil {
		panic(err)
	}

	defaultUser := model.User{
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "admin123",
		Roles:     []model.Role{roleAdministrator},
	}
	_, err = defaultUser.CreateUser()
	if err != nil {
		panic(err)
	}

	log.Println("Created default user with admin privileges... Please save these credentials!")
	log.Println("Email: admin@example.com")
	log.Println("Password: admin123")
}
