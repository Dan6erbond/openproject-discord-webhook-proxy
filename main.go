package main

import (
	"fmt"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/public"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	app := public.NewApp()
	app.Run()
}
