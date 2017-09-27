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
	Key   string
	Token string
	Board string
}

var config Config

func main() {
	fConfig := flag.String("c", "./config.toml", "config file path")
	flag.Parse()

	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		log.Fatalln(err)
	}

	trello, err := trello.NewAuthClient(config.Trello.Key, &config.Trello.Token)
	if err != nil {
		log.Fatalln(err)
	}

	board, err := trello.Board(config.Trello.Board)
	if err != nil {
		log.Fatalln(err)
	}

	cards, err := board.Cards()
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range cards {
		fmt.Println(v.Name)
	}
}
