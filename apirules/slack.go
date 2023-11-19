package apirules

import (
	"fmt"
	"github.com/slack-go/slack"
)

/*
let token = "";
let channel = "C066S807J2D";
let title = "DCJ";
let message = `new message ${RQL.RandInt(1, 100)}`;
let colour = "#FF0000";

RQL.Result = RQL.Slack(token, channel, title, message, colour);
*/

func (inst *RQL) Slack(token, channelId, title, message, colour string) any {
	api := slack.New(token)
	if colour == "" {
		colour = "#3AA3E3"
	}
	attachment := slack.Attachment{
		Title: title,
		Text:  message,
		Color: colour,
	}
	channelID, _, err := api.PostMessage(
		channelId,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		return err
	}
	return fmt.Sprintf("message successfully sent to channel %s", channelID)
}
