package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"github.com/phn00dev/go-blog-temp/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRouter()
	routing.Serve()
}
