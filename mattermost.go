package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func LoginAsTheBotUser() {

	client.SetToken(viper.GetString("bot_token"))

	if user, resp := client.GetUser(viper.GetString("bot_id"), ""); resp.Error != nil {
		log.WithError(resp.Error).Fatal("There was a problem getting Bot User")
	} else {
		botUser = user
	}

	// If you want to login via username/password
	//
	// if user, resp := client.Login(USER_EMAIL, USER_PASSWORD); resp.Error != nil {
	//  log.WithError(resp.Error).Fatal("There was a problem logging into the Mattermost server")
	// } else {
	// 	botUser = user
	// }
}

func MakeSureServerIsRunning() {
	if props, resp := client.GetOldClientConfig(""); resp.Error != nil {
		log.Fatal("There was a problem pinging the Mattermost server. Are you sure it's running?")

	} else {
		log.WithField("version", props["Version"]).Info("Server detected and is running")
	}
}

func FindBotTeam() {
	if team, resp := client.GetTeamByName(viper.GetString("team"), ""); resp.Error != nil {
		log.WithField("team", viper.GetString("team")).Fatal("Failed to get the initial load or we do not appear to be a member of the team")
	} else {
		botTeam = team
	}
}
func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			if webSocketClient != nil {
				webSocketClient.Close()
			}
			log.Info("Shutdown complete")
			os.Exit(0)
		}
	}()
}
