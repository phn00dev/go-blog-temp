package html

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/internal/providers/view"
)

func Render(ctx *gin.Context, statusCode int, pathName string, data gin.H) {
	data = view.WithGlobalData(data)
	ctx.HTML(statusCode, pathName, data)
}
