package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "github.com/phn00dev/go-blog-temp/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
