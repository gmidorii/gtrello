package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func slackSend(token, channel, text string) error {
	client := slack.New(token)
	message := fmt.Sprintf("```\n%s\n```", text)
	_, _, err := client.PostMessage(channel, message, slack.PostMessageParameters{
		Username: "Daily Report",
		Markdown: true,
	})
	return err
}
