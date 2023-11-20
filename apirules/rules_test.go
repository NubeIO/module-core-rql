package apirules

import (
	"fmt"
	"testing"

	"github.com/slack-go/slack"
)

func TestPG(t *testing.T) {
	var ChannelId = "C066S807J2D"
	api := slack.New("")
	attachment := slack.Attachment{
		Title: "DCJ",
		Text:  "count: 66 <@Aidan>",
		Color: "#FF0000",
	}

	users, err := api.GetUsers()
	if err != nil {
		return
	}

	for _, user := range users {
		fmt.Println(user.ID)
		fmt.Println(user.Name)
	}

	//fmt.Println(users)

	channelID, timestamp, err := api.PostMessage(
		ChannelId,
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		//slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}
