package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/database"
	"github.com/phn00dev/go-blog-temp/pkg/html"
	"github.com/phn00dev/go-blog-temp/pkg/routing"
	"github.com/phn00dev/go-blog-temp/pkg/sessions"
	"github.com/phn00dev/go-blog-temp/pkg/static"
)

func Serve() {
	config.Set()
	// database connect

	database.Connect()

	routing.Init()

	getRouter := routing.GetRouter()
	sessions.Start(getRouter)
	static.LoadStatic(getRouter)
	html.LoadHTML(getRouter)
	routing.RegisterRouter()
	routing.Serve()
}
