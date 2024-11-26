package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"net/http"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) HomePage(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page ",
	})
}
