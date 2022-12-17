package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	Host string

	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func InitConfig() (*Config, error) {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error read config env file: %w", err)
	}

	return &Config{
		Port:       viper.Get("SERVER_PORT").(string),
		Host:       viper.Get("SERVER_HOST").(string),
		DBUser:     viper.Get("POSTGRES_USER").(string),
		DBPassword: viper.Get("POSTGRES_PASSWORD").(string),
		DBName:     viper.Get("POSTGRES_DB").(string),
		DBHost:     viper.Get("POSTGRES_HOST").(string),
		DBPort:     viper.Get("POSTGRES_PORT").(string),
	}, nil
}
