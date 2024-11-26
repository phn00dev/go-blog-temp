package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config yaml file error")
		return
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct, %v", err)
		return
	}
}
