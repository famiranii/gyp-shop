package config

import "github.com/spf13/viper"

type Config struct {
	Port string
}

func LoadConfig() Config {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return Config{
		Port: viper.GetString("PORT"),
	}
}