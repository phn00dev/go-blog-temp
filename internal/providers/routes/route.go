package routes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "github.com/phn00dev/go-blog-temp/internal/modules/acticle/routes"
	homeRoutes "github.com/phn00dev/go-blog-temp/internal/modules/home/routes"
	userRoutes "github.com/phn00dev/go-blog-temp/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.UserRoutes(router)
}
