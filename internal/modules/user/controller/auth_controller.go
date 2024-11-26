package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/dto/authDTO"
	"github.com/phn00dev/go-blog-temp/internal/modules/user/service"
	"github.com/phn00dev/go-blog-temp/pkg/converters"
	"github.com/phn00dev/go-blog-temp/pkg/errors"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"github.com/phn00dev/go-blog-temp/pkg/old"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"log"
	"net/http"
	"strconv"
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
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	if a.userService.CheckUserExists(registerRequest.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	// create user
	user, err := a.userService.Create(registerRequest)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/register")
	}
	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))
	log.Printf("user created successfully username: %v \n", user.Name)
	ctx.Redirect(http.StatusFound, "/")
}

func (a *AuthController) Login(ctx *gin.Context) {
	html.Render(ctx, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (a *AuthController) LoginHandler(ctx *gin.Context) {
	var loginRequest authDTO.LoginRequest
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		errors.Init()
		errors.SetFormErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := a.userService.HandleUserLogin(loginRequest)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))
		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))
	log.Printf("user logged in successfully username: %v \n", user.Name)
	ctx.Redirect(http.StatusFound, "/")
}

func (a *AuthController) Logout(ctx *gin.Context) {
	sessions.Remove(ctx, "auth")
	ctx.Redirect(http.StatusFound, "/")
}
