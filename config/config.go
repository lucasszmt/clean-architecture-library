package config

import (
	"github.com/spf13/viper"
	"log"
)

var viperConfig *viper.Viper

type config struct {
	ApiHost     string            `mapstructure:"api_host"`
	ApiPort     string            `mapstructure:"api_port"`
	TesteNested map[string]string `mapstructure:"teste"`
}

func Init() {
	viperConfig = viper.New()
	viperConfig.SetConfigName("env-development")
	viperConfig.SetConfigType("yaml")
	viperConfig.AddConfigPath("./config")
	readErr := viperConfig.ReadInConfig()
	if readErr != nil {
		log.Fatal(readErr)
	}
}

func GetConfig() *config {
	var c *config
	err := viperConfig.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
