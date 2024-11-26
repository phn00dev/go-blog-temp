package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/modules/home/controller"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"net/http"
)

func Routes(router *gin.Engine) {
	homeController := controller.New()
	router.GET("/", homeController.HomePage)

	router.GET("/about", func(ctx *gin.Context) {
		html.Render(ctx, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page ",
		})
	})

}
