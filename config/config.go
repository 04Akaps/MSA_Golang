package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	MongoAddress  string `mapstructure:"mongo_address"`
}

func LoadConfig(path string) Config {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("env Read Error : &w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("env Marshal Error : &w", err)
	}

	return config
}
