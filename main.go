package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/BurntSushi/toml"
	trello "github.com/VojtechVitek/go-trello"
	"github.com/pkg/errors"
)

type Flag struct {
	Config   string
	Template string
	Output   string
}

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

func main() {
	myFlag := parseFlag()
	var config Config
	if _, err := toml.DecodeFile(myFlag.Config, &config); err != nil {
		log.Fatalf("failed config file :%+v\n", err)
	}

	output, err := pullTodo(config)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	name, err := writeFile(myFlag.Template, output, myFlag.Output)
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

	var a string
	for {
		fmt.Print("slack post ? y/n: ")
		fmt.Scan(&a)
		if a == "y" || a == "n" {
			break
		}
		fmt.Println("ERROR: input permit 'y' or 'n'")
	}
	if a == "n" {
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = slackSend(config.Slack.Token, config.Slack.Channel, string(b))
	if err != nil {
		log.Fatalf("failed send :%+v\n", err)
	}
}

func parseFlag() Flag {
	myFlag := Flag{
		Config:   *flag.String("c", "./config.toml", "config file path"),
		Template: *flag.String("t", "./template/template.md", "template file path"),
		Output:   *flag.String("o", "./", "output file path"),
	}
	flag.Parse()

	return myFlag
}

func pullTodo(config Config) (Output, error) {
	var output Output
	client, err := trello.NewAuthClient(config.Trello.Key, &config.Trello.Token)
	if err != nil {
		return output, errors.Wrap(err, "failed auth trello")
	}

	s := time.Now()
	output, err = fetchTrello(config.Trello.BoardID, client)
	if err != nil {
		return output, errors.Wrap(err, "failed fetch todo")
	}
	fmt.Printf("%f s\n", time.Now().Sub(s).Seconds())

	return output, nil
}
