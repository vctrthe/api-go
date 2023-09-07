package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	JWT      JwtConfig
	Midtrans MidtransConfig
}

type DatabaseConfig struct {
	Driver    string
	Username  string
	Password  string
	Host      string
	Port      string
	DBName    string
	Charset   string
	ParseTime string
	Loc       string
}

type JwtConfig struct {
	Secret string
}

type MidtransConfig struct {
	ServerKey string
	ClientKey string
}

var C Config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}
}
