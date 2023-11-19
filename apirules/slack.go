package apirules

import (
	"fmt"
	"github.com/slack-go/slack"
)

func (inst *RQL) Slack(token, channelId, message string) any {
	api := slack.New(token)
	channelID, _, err := api.PostMessage(
		channelId,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		return err
	}
	return fmt.Sprintf("message successfully sent to channel %s", channelID)
}
