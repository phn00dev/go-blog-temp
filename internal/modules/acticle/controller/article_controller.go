package controller

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/phn00dev/go-blog-temp/internal/modules/acticle/service"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"net/http"
	"strconv"
)

type ArticleController struct {
	articleService articleService.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: articleService.NewArticleServiceImp(),
	}
}

func (a *ArticleController) Show(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		html.Render(ctx, http.StatusBadRequest, "templates/errors/html/500", gin.H{
			"status_code": http.StatusBadRequest,
			"msg":         "Bad Request",
		})
		return
	}
	// find article
	articleResponse, err := a.articleService.Find(id)
	if err != nil {
		html.Render(ctx, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"status_code": http.StatusNotFound,
			"msg":         "Page not found",
			"errors":      err.Error(),
		})
		return
	}

	html.Render(ctx, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "show article",
		"article": articleResponse,
	})

}

func (a *ArticleController) Create(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}
