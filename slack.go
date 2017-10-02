package main

import "github.com/nlopes/slack"

func send(token string) error {
	api := slack.New(token)
}
