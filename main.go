// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/spf13/viper"
)

var client *model.Client4
var webSocketClient *model.WebSocketClient

var botUser *model.User

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	SetupGracefulShutdown()
	readConfig()
	writeConfig()

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	client = model.NewAPIv4Client(viper.GetString("api_server"))

	// Lets test to see if the mattermost server is up and running
	MakeSureServerIsRunning()

	// lets attempt to login to the Mattermost server as the bot user
	LoginAsTheBotUser()

	// Lets start listening to some channels via the websocket!
	webSocketClient, err := model.NewWebSocketClient4(viper.GetString("ws_server"), client.AuthToken)
	if err != nil {
		log.WithError(err).Fatal("We failed to connect to the web socket")
	}

	webSocketClient.Listen()

	go func() {
		for resp := range webSocketClient.EventChannel {
			HandleWebSocketResponse(resp)
		}
	}()

	// You can block forever with
	select {}
}

func HandleWebSocketResponse(event *model.WebSocketEvent) {
	HandleMsgFromStonksChannel(event)
}
