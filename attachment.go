package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

var colors = []string{
	"#078E38",
	"#EB9E5A",
	"#4EC4E7",
	"#914EE7",
	"9CA0AF",
}

func CreateAttachements(todo Todo) []slack.Attachment {
	attachments := make([]slack.Attachment, len(todo.Lists))
	for i, list := range todo.Lists {
		var color string
		if i < len(colors) {
			color = colors[i]
		}
		attachments[i] = createAttachement(list, color)
	}
	return attachments
}

func createAttachement(list TodoList, color string) slack.Attachment {
	var text string
	for _, card := range list.Cards {
		text += fmt.Sprintf("・%s\n", card.Name)
	}
	return slack.Attachment{
		Title: list.Name,
		Text:  text,
		Color: color,
	}
}
