package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DataBaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Server   ServerConfig
	DataBase DataBaseConfig
	PageSize int `yaml:"PageSize"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable decode into struct, %v", err)
	}

	return &cfg
}
