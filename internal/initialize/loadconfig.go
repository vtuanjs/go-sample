package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/vtuanjs/my-project/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %v", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}
}
