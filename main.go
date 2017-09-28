package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	trello "github.com/VojtechVitek/go-trello"
)

type Config struct {
	Trello Trello
}

type Trello struct {
	Key     string
	Token   string
	BoardID string
}

const gtrello = "gtrello.md"

var (
	config       Config
	templateFile string
)

func main() {
	fConfig := flag.String("c", "./config.toml", "config file path")
	fTemplate := flag.String("t", ".template/template.md", "template file path")
	flag.Parse()

	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		log.Fatalln(err)
	}

	client, err := trello.NewAuthClient(config.Trello.Key, &config.Trello.Token)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := createList(config.Trello.BoardID, client)
	if err != nil {
		log.Fatalln(err)
	}
	if err = outputFile(*fTemplate, output, gtrello); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("vim", gtrello)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
