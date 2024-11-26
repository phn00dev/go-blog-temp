package routes

import (
	"github.com/gin-gonic/gin"
	authController "github.com/phn00dev/go-blog-temp/internal/modules/user/controller"
)

func UserRoutes(router *gin.Engine) {
	authController := authController.NewAuthController()
	// register routes
	router.GET("/register", authController.Register)
	router.POST("/register", authController.RegisterHandler)

	// login routes

	router.GET("/login", authController.Login)
	router.POST("/login", authController.LoginHandler)

	router.POST("/logout", authController.Logout)

}
