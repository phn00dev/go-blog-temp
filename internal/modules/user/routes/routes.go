package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/middlewares"
	authCtrl "github.com/phn00dev/go-blog-temp/internal/modules/user/controller"
)

func UserRoutes(router *gin.Engine) {
	authController := authCtrl.NewAuthController()
	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		// register routes
		guestGroup.GET("/register", authController.Register)
		guestGroup.POST("/register", authController.RegisterHandler)
		// login routes
		guestGroup.GET("/login", authController.Login)
		guestGroup.POST("/login", authController.LoginHandler)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", authController.Logout)
	}

}
