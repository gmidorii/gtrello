package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
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

	todo, err := PullTodo(config.Trello)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	name, err := writeFile(myFlag.Template, todo, myFlag.Output)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err = editVim(name); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if isPostSlack() {
		err = postSlack(name, config.Slack)
		if err != nil {
			log.Fatalln(err)
		}
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

func editVim(name string) error {
	cmd := exec.Command("vim", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "edit failed")
	}
	return nil
}

func isPostSlack() bool {
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
		return false
	}

	return true
}

func postSlack(name string, slack Slack) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = slackSend(slack.Token, slack.Channel, string(b))
	if err != nil {
		return errors.Wrap(err, "failed send")
	}
	return nil
}
