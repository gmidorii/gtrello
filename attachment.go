package main

import (
	"fmt"
	"io/ioutil"

	"github.com/nlopes/slack"
)

var colors = []string{
	"#078E38",
	"#EB9E5A",
	"#4EC4E7",
	"#914EE7",
	"9CA0AF",
}

func CreateAttachements(todo Todo, outputFile string, daylists []string) ([]slack.Attachment, error) {
	attachments := make([]slack.Attachment, len(todo.Lists)+1)
	for i, list := range todo.Lists {
		if !containsDaylists(list.Name, daylists) {
			continue
		}
		var color string
		if i < len(colors) {
			color = colors[i]
		}
		attachments[i] = createAttachement(list, color)
	}

	o, err := ioutil.ReadFile(outputFile)
	if err != nil {
		return nil, err
	}
	// Fix later
	attachments[len(attachments)-1] = slack.Attachment{
		Title: "Comment",
		Text:  string(o),
		Color: "#066f6f",
	}

	return attachments, nil
}

func createAttachement(list TodoList, color string) slack.Attachment {
	var text string
	for _, card := range list.Cards {
		text += fmt.Sprintf("■  %s\n", card.Name)
	}
	return slack.Attachment{
		Title: list.Name,
		Text:  text,
		Color: color,
	}
}

func containsDaylists(name string, daylists []string) bool {
	for _, day := range daylists {
		if day == name {
			return true
		}
	}
	return false
}
