package config

import (
	"github.com/spf13/viper"
	"log"
)

func GetConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}