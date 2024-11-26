package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/providers/routes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRouter() {
	routes.RegisterRoutes(GetRouter())
}
