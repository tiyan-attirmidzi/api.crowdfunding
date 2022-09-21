package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tiyan-attirmidzi/api.crowdfunding/controllers"
	"github.com/tiyan-attirmidzi/api.crowdfunding/repositories"
	"github.com/tiyan-attirmidzi/api.crowdfunding/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// DB CONFIGURATION
	dsn := "ro0t:P@5sw012D@tcp(127.0.0.1:3306)/learn_api_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Database Connected Successfully!")

	// REPOSITORY, SERVICE, CONTROLLER and ROUTE
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(authService, userService)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			auth.POST("/sign-up", authController.SignUp)
			auth.POST("/sign-in", authController.SignIn)
		}
		user := apiV1.Group("/users")
		{
			user.POST("/avatar", userController.UploadAvatar)
			user.POST("/email-check", userController.CheckEmailAvailability)
		}
	}

	router.Run()

}
