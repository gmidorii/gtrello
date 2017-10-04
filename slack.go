package main

import (
	"errors"
	"fmt"

	"github.com/nlopes/slack"
)

func send(token string) error {
	api := slack.New(token)
	fmt.Println(api)
	return errors.New("err")
}
