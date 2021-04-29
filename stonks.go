package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/spf13/viper"
)

func HandleMsgFromStonksChannel(event *model.WebSocketEvent) {

	// Lets only respond to posted events
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	channel, resp := client.GetChannel(event.Broadcast.ChannelId, "")
	if resp.Error != nil {
		log.WithField("channel_id", event.Broadcast.ChannelId).WithError(resp.Error).Error("Failed to get channel by ID")
		return
	}

	re := regexp.MustCompile(`^ \*\*Deal won by(?:$|\W)`)

	if Contains(viper.GetStringSlice("stonks.channels"), channel.Name) {
		log.WithField("channel_id", channel.Name).Debug("Channel is contained in stringSlice from config")

		post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
		if post == nil {
			log.Error("Failed to read post")
			return
		}

		matched := re.MatchString(post.Message)
		if matched {
			go func() {
				addStonksReaction(post, channel.Name)
			}()
		} else {
			log.WithField("message", post.Message).Debug("Message did not match the regex")
		}
	}
}

func addStonksReaction(post *model.Post, channel string) {

	min := viper.GetInt("stonks.mindelay")
	max := viper.GetInt("stonks.maxdelay")

	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(max-min+1) + min

	log.WithFields(log.Fields{
		"channel": channel,
		"message": post.Message,
		"emojis":  viper.GetStringSlice("stonks.emojis"),
		"delay":   delay,
	}).Info("Message which matches the regex. Adding reactions after delay")

	time.Sleep(time.Duration(delay) * time.Second)

	for _, emoji := range viper.GetStringSlice("stonks.emojis") {
		reaction := &model.Reaction{
			UserId:    botUser.Id,
			PostId:    post.Id,
			EmojiName: emoji,
		}
		if _, resp := client.SaveReaction(reaction); resp.Error != nil {
			log.WithError(resp.Error).Error("Failed to add reaction to post")
		}

	}
}
