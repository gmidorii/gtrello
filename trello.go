package main

import (
	"fmt"
	"time"

	trello "github.com/VojtechVitek/go-trello"
	"github.com/pkg/errors"
)

const (
	trelloLayout = "2006-01-02T15:04:05.000Z"
	outputLayout = "01/02"
)

type Output struct {
	Lists []TmpList
}

type TmpList struct {
	ID    string
	Name  string
	Cards []TmpCard
}

type TmpCard struct {
	Name       string
	DeadLine   string
	Checklists []TmpCheckList
}

type TmpCheckList struct {
	Name       string
	CheckItems []TmpCheckItem
}

type TmpCheckItem struct {
	Name  string
	State string
}

func fetchTrello(boardID string, client *trello.Client) (Output, error) {
	var output Output

	start := time.Now()
	board, err := client.Board(boardID)
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello board")
	}
	fmt.Printf("board: %g s\n", time.Now().Sub(start).Seconds())

	// lists
	start = time.Now()
	lists, err := board.Lists()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello lists")
	}
	fmt.Printf("lists: %g s\n", time.Now().Sub(start).Seconds())
	var tmpLists []TmpList
	for _, v := range lists {
		tmpLists = append(tmpLists, TmpList{
			ID:   v.Id,
			Name: v.Name,
		})
	}

	start = time.Now()
	cards, err := board.Cards()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello cards")
	}
	fmt.Printf("cards: %g s\n", time.Now().Sub(start).Seconds())

	start = time.Now()
	for _, card := range cards {
		for i, list := range tmpLists {
			if card.IdList == list.ID {
				checklists, err := card.Checklists()
				if err != nil {
					return output, errors.Wrap(err, "faild fetch trello checklist from card")
				}
				var tmpCheckLists []TmpCheckList
				for _, checklist := range checklists {
					var tmpItems []TmpCheckItem
					for _, item := range checklist.CheckItems {
						tmpItems = append(tmpItems, TmpCheckItem{
							Name:  item.Name,
							State: item.State,
						})
					}
					tmpCheckLists = append(tmpCheckLists, TmpCheckList{
						Name:       checklist.Name,
						CheckItems: tmpItems,
					})
				}
				var dstr string
				if card.Due != "" {
					deadLine, err := time.Parse(trelloLayout, card.Due)
					if err != nil {
						return output, errors.Wrap(err, "faild due parse")
					}
					dstr = deadLine.Format(outputLayout)
				}
				tmpLists[i].Cards = append(tmpLists[i].Cards, TmpCard{
					Name:       card.Name,
					DeadLine:   dstr,
					Checklists: tmpCheckLists,
				})
				continue
			}
		}
	}
	fmt.Printf("convert: %g s\n", time.Now().Sub(start).Seconds())

	return Output{tmpLists}, nil
}
