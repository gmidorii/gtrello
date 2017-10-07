package main

import (
	"errors"

	"github.com/nlopes/slack"
)

func send(token, channel, text string) error {
	client := slack.New(token)
	client.PostMessage(channel, text, slack.PostMessageParameters{
		Username: "Daily Report",
		Markdown: true,
	})
	return errors.New("err")
}
