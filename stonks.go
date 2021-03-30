package main

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/spf13/viper"
)

func HandleMsgFromStonksChannel(event *model.WebSocketEvent) {

	// Lets only reponded to messaged posted events
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	re := regexp.MustCompile(`^ \*\*Deal won by(?:$|\W)`)

	for _, allowedChannel := range viper.GetStringSlice("stonks.channels") {
		channel, resp := client.GetChannelByName(allowedChannel, botTeam.Id, "")
		if resp.Error != nil {

			log.WithFields(log.Fields{
				"channel": allowedChannel,
				"error":   resp.Error,
			}).Error("Failed to get channel")

			continue
		}

		logger := log.WithFields(log.Fields{
			"channel": allowedChannel,
		})

		if event.Broadcast.ChannelId == channel.Id {
			post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
			if post == nil {
				logger.Error("Failed to read post")
				continue
			}

			matched := re.MatchString(post.Message)
			if matched {

				for _, emoji := range viper.GetStringSlice("stonks.emojis") {
					reaction := &model.Reaction{
						UserId:    botUser.Id,
						PostId:    post.Id,
						EmojiName: emoji,
					}
					// TODO: implement delay for reaction based on viper.GetInt("stonks.maxdelay")
					if _, resp := client.SaveReaction(reaction); resp.Error != nil {
						log.WithError(resp.Error).Error("Failed to add reaction to post")
					}

				}
				logger.WithFields(log.Fields{
					"message": post.Message,
					"emojis":  viper.GetStringSlice("stonks.emojis"),
				}).Info("Message which matches the regex. Added reactions")

			}
		}
	}
}
