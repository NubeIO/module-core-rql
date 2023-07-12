package apirules

import (
	"fmt"
	"testing"

	"github.com/slack-go/slack"
)

func TestPG(t *testing.T) {
	var ChannelId = "C05DNBFP1M4"
	api := slack.New("")
	//attachment := slack.Attachment{
	//	Pretext: "alert",
	//	Text:    "<@UJ6T8ALCR> <@aidan> alert from device ABC",
	//	// Uncomment the following part to send a field too
	//	/*
	//		Fields: []slack.AttachmentField{
	//			slack.AttachmentField{
	//				Title: "a",
	//				Value: "no",
	//			},
	//		},
	//	*/
	//}

	channelID, timestamp, err := api.PostMessage(
		ChannelId,
		slack.MsgOptionText("Ping failed <@UJ6T8ALCR>", false),
		//slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}
