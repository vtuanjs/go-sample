package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	App struct {
		Name string
		Port int
	}
}

func main() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()

	fmt.Println(viper.Get("app.name"))

	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
