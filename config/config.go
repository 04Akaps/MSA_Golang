package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	MongoAddress  string `mapstructure:"mongo_address"`
	MongoHostName string `mapstructure:"mongo_host_name"`
	MongoPort     string `mapstructure:"mongo_port"`
	MongoUserName string `mapstructure:"mongo_user_name"`
	MongoPassword string `mapstructure:"mongo_password"`
	CsrName       string `mapstructure:"csr_name"`
	KeyName       string `mapstructure:"key_name"`
	AmqpUri       string `mapstructure:"amqp_uri"`
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
