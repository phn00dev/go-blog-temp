package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	routing.RegisterRouter()
	routing.Serve()
}
