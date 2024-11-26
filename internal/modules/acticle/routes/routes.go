package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/middlewares"
	"github.com/phn00dev/go-blog-temp/internal/modules/acticle/controller"
)

func Routes(router *gin.Engine) {
	articleController := controller.NewArticleController()
	router.GET("/articles/:ID", articleController.Show)
	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/create", articleController.Store)
	}

}
