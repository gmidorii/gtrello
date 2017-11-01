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
		Title:    "Week Report",
		Channels: []string{channel},
	})
	if err != nil {
		return errors.Wrap(err, "faild post channel")
	}
	return nil
}

func slackSendAttachment(token, channel string, attachments []slack.Attachment) error {
	client := slack.New(token)
	_, _, err := client.PostMessage(channel, "", slack.PostMessageParameters{
		Attachments: attachments,
	})
	if err != nil {
		return errors.Wrap(err, "faild post attachments channel")
	}
	return nil
}
