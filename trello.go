package main

import (
	trello "github.com/VojtechVitek/go-trello"
	"github.com/pkg/errors"
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

func featchTrello(boardID string, client *trello.Client) (Output, error) {
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
	var tmpLists []TmpList
	for _, v := range lists {
		tmpLists = append(tmpLists, TmpList{
			ID:   v.Id,
			Name: v.Name,
		})
	}

	cards, err := board.Cards()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello cards")
	}

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
				tmpLists[i].Cards = append(tmpLists[i].Cards, TmpCard{
					Name:       card.Name,
					Checklists: tmpCheckLists,
				})
				continue
			}
		}
	}

	return Output{tmpLists}, nil
}
