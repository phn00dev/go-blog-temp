package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-blog-temp/config"
	"github.com/spf13/viper"
)

func main() {
	configs := configSet()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app name": viper.Get("app.AppName"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func configSet() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("read config yaml file error")
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	return configs
}
