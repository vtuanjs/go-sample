package initialize

import (
	"fmt"
	"vtuanjs/my-project/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %v", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %v", err))
	}
}
