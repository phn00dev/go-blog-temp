package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"github.com/phn00dev/go-blog-temp/pkg/routing"
	"github.com/phn00dev/go-blog-temp/pkg/static"
)

func Serve() {
	config.Set()
	routing.Init()
	getRouter := routing.GetRouter()
	static.LoadStatic(getRouter)
	html.LoadHTML(getRouter)
	routing.RegisterRouter()
	routing.Serve()
}
