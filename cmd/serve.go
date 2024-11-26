package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/routing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run project ",
	Long:  `run project`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config.Set()
	routing.Init()
	router := routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app name": viper.Get("app.AppName"),
		})
	})
	routing.Serve()
}
