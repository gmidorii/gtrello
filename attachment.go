package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func CreateAttachements(todo Todo) []slack.Attachment {
	attachments := make([]slack.Attachment, len(todo.Lists))
	for i, list := range todo.Lists {
		attachments[i] = createAttachement(list)
	}
	return attachments
}

func createAttachement(list TodoList) slack.Attachment {
	var text string
	for _, card := range list.Cards {
		text += fmt.Sprintf("%s\n", card.Name)
	}
	return slack.Attachment{
		Title: list.Name,
		Text:  text,
	}
}
