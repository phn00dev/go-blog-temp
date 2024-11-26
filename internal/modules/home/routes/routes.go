package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page ",
		})
	})
	router.GET("/about", func(ctx *gin.Context) {
		html.Render(ctx, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page ",
		})
	})

}
