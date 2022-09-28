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
		"Low":       "🔻",
		"Normal":    "🔲",
		"High":      "🔺",
		"Immediate": "❗",
	})
	viper.SetDefault("openproject.statusEmojiMappings", map[string]string{
		"New":              "💠",
		"In specification": "🔜",
		"Specified":        "🟢",
		"Confirmed":        "✔️",
		"To be  scheduled": "🕐",
		"Scheduled":        "🕐",
		"In progress":      "⏩",
		"Developed":        "💯",
		"In testing":       "⏩",
		"Tested":           "💯",
		"Test failed":      "⚠️",
		"Closed":           "⛔",
		"On hold":          "⏸️",
		"Rejected":         "❌",
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	app := pkg.NewApp()
	app.Run()
}
