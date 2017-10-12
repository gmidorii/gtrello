package main

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

func slackSend(token, channel, text string) error {
	client := slack.New(token)
	_, err := client.UploadFile(slack.FileUploadParameters{
		Content:  text,
		Filetype: "markdown",
		Title:    "Daily",
		Channels: []string{channel},
	})
	if err != nil {
		return errors.Wrap(err, "faild post channel")
	}
	return nil
}
