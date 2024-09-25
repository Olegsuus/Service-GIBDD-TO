package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DataBaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type Config struct {
	Server   ServerConfig
	DataBase DataBaseConfig
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.BindEnv("Server.Port", "SERVER_PORT")
	viper.BindEnv("DataBase.Driver", "DB_DRIVER")
	viper.BindEnv("DataBase.Host", "DB_HOST")
	viper.BindEnv("DataBase.Port", "DB_PORT")
	viper.BindEnv("DataBase.User", "DB_USER")
	viper.BindEnv("DataBase.Password", "DB_PASSWORD")
	viper.BindEnv("DataBase.DBName", "DB_NAME")

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
