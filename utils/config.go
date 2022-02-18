package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Config struct {
	Database []DBConfig     `mapstructure:"database"`
	Server   []ServerConfig `mapstructure:"server"`
}

func LoadConfig() Config {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Print(err.Error())

		return config
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Print(err.Error())
		return config
	}

	return config
}
