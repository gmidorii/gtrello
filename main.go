package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	trello "github.com/VojtechVitek/go-trello"
)

const trelloUrl = "https://trello.com/1"

type Config struct {
	Trello Trello
}

type Trello struct {
	Key     string
	Token   string
	BoardID string
}

var config Config

func main() {
	fConfig := flag.String("c", "./config.toml", "config file path")
	flag.Parse()

	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		log.Fatalln(err)
	}

	client, err := trello.NewAuthClient(config.Trello.Key, &config.Trello.Token)
	if err != nil {
		log.Fatalln(err)
	}

	lists, err := createList(config.Trello.BoardID, client)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lists)
}
