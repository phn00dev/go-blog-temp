package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/dto/authDTO"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/service"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"log"
	"net/http"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: service.NewUserService(),
	}
}

func (a *AuthController) Register(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (a *AuthController) RegisterHandler(ctx *gin.Context) {
	var registerRequest authDTO.RegisterRequest
	if err := ctx.ShouldBind(&registerRequest); err != nil {
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	// create user
	user, err := a.userService.Create(registerRequest)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/register")
	}

	log.Printf("user created successfully username: %v \n", user.Name)
	ctx.Redirect(http.StatusFound, "/")
}

func (a *AuthController) Login(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (a *AuthController) LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title": "Login done",
	})
}
