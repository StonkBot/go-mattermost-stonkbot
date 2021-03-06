package main

//TODO: Use https://github.com/spf13/cobra to have a fancy cmd interface to the bot

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	// "github.com/knadh/koanf"
)

func readConfig() {

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/stonks")
	viper.AddConfigPath("$HOME/.config/stonks")
	viper.AddConfigPath(".")

	viper.SetDefault("debug", false)

	viper.SetDefault("api_server", "http://localhost:8065")
	viper.SetDefault("ws_server", "ws://localhost:8065")
	viper.SetDefault("team", "test")

	viper.SetDefault("bot_id", "")
	viper.SetDefault("bot_token", "")

	viper.SetDefault("stonks.channels", [0]string{})
	viper.SetDefault("stonks.emojis", [1]string{"stonks"})
	viper.SetDefault("stonks.mindelay", 10)
	viper.SetDefault("stonks.maxdelay", 30)

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Info("Failed to read configfile. Saving an empty one")
		viper.SafeWriteConfig()
	}

	if viper.GetString("bot_id") == "" || viper.GetString("bot_token") == "" {
		log.Fatal("No bot id or token configured")
	}

	if viper.GetInt("stonks.mindelay") > viper.GetInt("stonks.maxdelay") {
		log.Fatal("Stonks mindelay is greater than maxdelay")
	}

}

func writeConfig() {

	viper.WriteConfig()
}
