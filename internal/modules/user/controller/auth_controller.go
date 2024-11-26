package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"net/http"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Register(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (a *AuthController) RegisterHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title": "Register done",
	})
}

func (a *AuthController) Login(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (a *AuthController) LoginHandler(ctx *gin.Context) {}
