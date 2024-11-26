package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	articleService "github.com/phn00dev/go-blog-temp/internal/modules/acticle/service"
	"github.com/phn00dev/go-blog-temp/pkg/html"
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

	html.Render(ctx, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"stories":  controller.articleService.GetStories(6),
		"featured": controller.articleService.GetFeatured(4),
	})

}
