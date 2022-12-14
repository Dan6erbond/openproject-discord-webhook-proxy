package main

import (
	"fmt"

	"github.com/dan6erbond/openproject-discord-webhook-proxy/pkg"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", 5001)
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("openproject.priorityEmojiMappings", map[string]string{
		"Low":       "π»",
		"Normal":    "π²",
		"High":      "πΊ",
		"Immediate": "β",
	})
	viper.SetDefault("openproject.statusEmojiMappings", map[string]string{
		"New":              "π ",
		"In specification": "π",
		"Specified":        "π’",
		"Confirmed":        "βοΈ",
		"To be  scheduled": "π",
		"Scheduled":        "π",
		"In progress":      "β©",
		"Developed":        "π―",
		"In testing":       "β©",
		"Tested":           "π―",
		"Test failed":      "β οΈ",
		"Closed":           "β",
		"On hold":          "βΈοΈ",
		"Rejected":         "β",
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	app := pkg.NewApp()
	app.Run()
}
