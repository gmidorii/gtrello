package main

import (
	"github.com/nlopes/slack"
)

func slackSend(token, channel, text string) error {
	client := slack.New(token)
	_, _, err := client.PostMessage(channel, text, slack.PostMessageParameters{
		Text:     "text",
		Username: "Daily Report",
		Markdown: true,
	})
	return err
}
