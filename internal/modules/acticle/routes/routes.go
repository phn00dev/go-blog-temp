package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/modules/acticle/controller"
)

func Routes(router *gin.Engine) {
	articleController := controller.NewArticleController()
	router.GET("/articles/:ID", articleController.Show)
}
