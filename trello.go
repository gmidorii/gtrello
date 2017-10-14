package main

import (
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

	board, err := client.Board(boardID)
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello board")
	}

	// lists
	lists, err := board.Lists()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello lists")
	}
	tmpLists := make([]TmpList, len(lists), len(lists))
	for i, v := range lists {
		tmpLists[i] = TmpList{
			ID:   v.Id,
			Name: v.Name,
		}
	}

	cards, err := board.Cards()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello cards")
	}

	// fetch checklist
	checkChan := make(chan []trello.Checklist)
	for _, card := range cards {
		go func(card trello.Card) {
			checklist, _ := card.Checklists()
			checkChan <- checklist
		}(card)
	}
	checklists := make([]trello.Checklist, len(cards), len(cards))
	idx := 0
	count := 0
loop:
	for {
		select {
		case checklist := <-checkChan:
			for _, v := range checklist {
				checklists[idx] = v
				idx++
			}
			count++
			if len(cards) == count {
				break loop
			}
		}
	}

	for _, card := range cards {
		for i, list := range tmpLists {
			if card.IdList == list.ID {
				checklists := findCheckLists(checklists, card.IdCheckLists)
				if err != nil {
					return output, errors.Wrap(err, "faild fetch trello checklist from card")
				}
				var tmpCheckLists []TmpCheckList
				for _, checklist := range checklists {
					tmpItems := make([]TmpCheckItem, len(checklist.CheckItems), len(checklist.CheckItems))
					for i, item := range checklist.CheckItems {
						tmpItems[i] = TmpCheckItem{
							Name:  item.Name,
							State: item.State,
						}
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
	return Output{tmpLists}, nil
}

func findCheckLists(checklists []trello.Checklist, ids []string) []trello.Checklist {
	results := make([]trello.Checklist, 0, 0)
	for _, id := range ids {
		for _, checklist := range checklists {
			if id == checklist.Id {
				results = append(results, checklist)
			}
		}
	}
	return results
}
