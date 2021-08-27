package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
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
	dir, fileName := filepath.Split(path)
	viper.AddConfigPath(dir)
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func WriteConfig(path string, items map[string]string) error{
	filePath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	dir, _ := filepath.Split(filePath)
	viper.AddConfigPath(dir)
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	for key, value := range items {
		viper.Set(key ,value)
	}
	err = viper.WriteConfigAs(filePath)
	if err != nil {
		return err
	}
	return nil
}