package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
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

func ReadConfig(path string) {
	fileType := strings.Replace(filepath.Ext(path), ".", "",-1)
	dir, fileName := filepath.Split(path)
	fileName = strings.Replace(fileName, fileType, "", -1)
	viper.SetConfigType(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func WriteConfig(path string) {
	fileType := strings.Replace(filepath.Ext(path), ".", "",-1)
	dir, fileName := filepath.Split(path)
	fileName = strings.Replace(fileName, fileType, "", -1)
	viper.SetConfigType(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	viper.WriteConfigAs(path)
}