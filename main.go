package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	trello "github.com/VojtechVitek/go-trello"
)

type Config struct {
	Trello Trello
	Slack  Slack
}

type Trello struct {
	Key     string
	Token   string
	BoardID string
}

type Slack struct {
	Token   string
	Channel string
}

const gtrello = "gtrello.md"

var (
	config       Config
	templateFile string
)

func main() {
	fConfig := flag.String("c", "./config.toml", "config file path")
	fTemplate := flag.String("t", "./template/template.md", "template file path")
	fOutput := flag.String("o", "./", "output file path")
	flag.Parse()

	if _, err := toml.DecodeFile(*fConfig, &config); err != nil {
		log.Fatalf("failed config file :%+v\n", err)
	}

	client, err := trello.NewAuthClient(config.Trello.Key, &config.Trello.Token)
	if err != nil {
		log.Fatalf("failed auth trello :%+v\n", err)
	}

	output, err := fetchTrello(config.Trello.BoardID, client)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	name, err := writeFile(*fTemplate, output, *fOutput)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("vim", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = slackSend(config.Slack.Token, config.Slack.Channel, string(b))
	if err != nil {
		log.Fatalf("failed send :%+v\n", err)
	}
}
