package controller

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/phn00dev/go-blog-temp/internal/modules/acticle/service"
	"net/http"
)

type Controller struct {
	articleService articleService.ArticleService
}

func New() *Controller {
	return &Controller{
		articleService: articleService.NewArticleServiceImp(),
	}
}

func (controller *Controller) HomePage(ctx *gin.Context) {
	//html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page ",
	//})
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "home",
		"stories":  controller.articleService.GetStories(10),
		"featured": controller.articleService.GetFeatured(10),
	})
}
