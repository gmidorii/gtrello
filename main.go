package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Flag struct {
	Config   *string
	Template *string
	Output   *string
	Day      *bool
	Week     *bool
}

type Config struct {
	Trello Trello
	Slack  Slack
}

type Trello struct {
	Key       string
	Token     string
	BoardID   string
	Daylists  []string
	Weeklists []string
}

type Slack struct {
	Token   string
	Channel string
}

const (
	format = "2006-01-02"
)

func init() {
	cfgPath := filepath.Join(os.Getenv("HOME"), ".config", "gtrello")
	_, err := os.Stat(cfgPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(cfgPath, 0755)
		if err != nil {
			log.Fatalln(err)
		}
		os.Mkdir(filepath.Join(cfgPath, "output"), 0755)
		os.Mkdir(filepath.Join(cfgPath, "template"), 0755)
		copy(filepath.Join("template", "template.md"), filepath.Join(cfgPath, "template", "template.md"))

		c, err := os.Create(filepath.Join(cfgPath, "config.toml"))
		if err != nil {
			log.Fatalln(err)
		}
		defer c.Close()
		var config Config
		enc := toml.NewEncoder(c)
		err = enc.Encode(config)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func copy(in, out string) error {
	i, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(out, i, 0755)
}

func run() error {
	myFlag := parseFlag()
	var config Config
	if _, err := toml.DecodeFile(*myFlag.Config, &config); err != nil {
		return errors.Wrap(err, "failed config file :%+v\n")
	}

	todo, err := PullTodo(config.Trello)
	if err != nil {
		return err
	}

	if *myFlag.Day {
		err = postDayReport(myFlag, todo, config)
		if err != nil {
			return err
		}
	}

	if *myFlag.Week {
		err := postWeekReport(myFlag, todo, config)
		if err != nil {
			return err
		}
	}
	return nil
}

func postDayReport(myFlag Flag, todo Todo, config Config) error {
	now := time.Now()
	outputFile := fmt.Sprintf("%s/day/%s.md", *myFlag.Output, now.Format(format))

	if err := editVim(outputFile); err != nil {
		return err
	}
	attachs, err := CreateAttachements(todo, outputFile, config.Trello.Daylists)
	if err != nil {
		return err
	}
	err = slackSendAttachment(config.Slack.Token, config.Slack.Channel, attachs)
	if err != nil {
		return err
	}

	fmt.Println("Successful Post Attachment to Slack!!")
	return nil
}

func postWeekReport(myFlag Flag, todo Todo, config Config) error {
	now := time.Now()
	outputFile := fmt.Sprintf("%s/week/%s.md", *myFlag.Output, now.Format(format))

	if err := writeFile(*myFlag.Template, outputFile, todo, config.Trello.Weeklists); err != nil {
		return err
	}

	if err := appendFile(outputFile, *myFlag.Output+"/day", 5); err != nil {
		return err
	}

	if isPostSlack() {
		err := postSlack(outputFile, config.Slack)
		if err != nil {
			return err
		}
		fmt.Println("Successful Post to Slack!!")
	}
	return nil
}

func containsList(name string, list []string) bool {
	for _, v := range list {
		if name == v {
			return true
		}
	}
	return false
}

func parseFlag() Flag {
	cfgPath := filepath.Join(os.Getenv("HOME"), ".config", "gTrello")
	myFlag := Flag{
		Config:   flag.String("c", filepath.Join(cfgPath, "config.toml"), "config file path"),
		Template: flag.String("t", filepath.Join(cfgPath, "template", "template.md"), "template file path"),
		Output:   flag.String("o", filepath.Join(cfgPath, "output"), "output dir path"),
		Day:      flag.Bool("d", true, "day report"),
		Week:     flag.Bool("w", false, "week report"),
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

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
