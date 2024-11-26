package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/modules/acticle/dto"
	articleService "github.com/phn00dev/go-blog-temp/internal/modules/acticle/service"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/helpers"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/service"
	"github.com/phn00dev/go-blog-temp/pkg/converters"
	"github.com/phn00dev/go-blog-temp/pkg/errors"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"github.com/phn00dev/go-blog-temp/pkg/old"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"net/http"
	"strconv"
)

type ArticleController struct {
	articleService articleService.ArticleService
	userService    service.UserService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: articleService.NewArticleServiceImp(),
		userService:    service.NewUserService(),
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
		"article": articleResponse,
	})

}

func (a *ArticleController) Create(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (a *ArticleController) Store(ctx *gin.Context) {
	var createRequest dto.CreateArticleRequest
	if err := ctx.ShouldBind(&createRequest); err != nil {
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))
		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(ctx)

	article, err := a.articleService.Store(createRequest, user)
	if err != nil {
		ctx.Redirect(http.StatusFound, "articles/create")
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
