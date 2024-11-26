package view

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/pkg/converters"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"github.com/spf13/viper"
)

func WithGlobalData(ctx *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("app.AppName")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(ctx, "errors"))
	return data
}
