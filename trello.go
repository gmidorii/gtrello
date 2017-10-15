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

type Todo struct {
	Lists []TodoList
}

type TodoList struct {
	ID    string
	Name  string
	Cards []TodoCard
}

type TodoCard struct {
	Name       string
	DeadLine   string
	Checklists []TodoCheckList
}

type TodoCheckList struct {
	Name       string
	CheckItems []TodoCheckItem
}

type TodoCheckItem struct {
	Name  string
	State string
}

func PullTodo(ctrello Trello) (Todo, error) {
	var todo Todo
	client, err := trello.NewAuthClient(ctrello.Key, &ctrello.Token)
	if err != nil {
		return todo, errors.Wrap(err, "failed auth trello")
	}

	s := time.Now()
	todo, err = fetchTrello(ctrello.BoardID, client)
	if err != nil {
		return todo, errors.Wrap(err, "failed fetch todo")
	}
	fmt.Printf("%f s\n", time.Now().Sub(s).Seconds())

	return todo, nil
}

func fetchTrello(boardID string, client *trello.Client) (Todo, error) {
	var output Todo

	board, err := client.Board(boardID)
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello board")
	}

	lists, err := board.Lists()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello lists")
	}

	cards, err := board.Cards()
	if err != nil {
		return output, errors.Wrap(err, "faild fetch trello cards")
	}

	checklists := fetchCheckLists(cards)

	todo, err := convert(lists, cards, checklists)
	if err != nil {
		return Todo{}, errors.Wrap(err, "failed convert")
	}
	return todo, nil
}

func fetchCheckLists(cards []trello.Card) []trello.Checklist {
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
	return checklists
}

func convert(lists []trello.List, cards []trello.Card, checklists []trello.Checklist) (Todo, error) {
	var todo Todo
	todo.Lists = make([]TodoList, len(lists), len(lists))
	for i, v := range lists {
		todo.Lists[i] = TodoList{
			ID:   v.Id,
			Name: v.Name,
		}
	}

	for _, card := range cards {
		convertTodoList(card, checklists, todo.Lists)
	}
	return todo, nil
}

func convertTodoList(card trello.Card, checklists []trello.Checklist, todoList []TodoList) error {
	for i, list := range todoList {
		if card.IdList == list.ID {
			checklists := findCheckLists(checklists, card.IdCheckLists)
			todoCheckLists := convertCheckLists(checklists)

			var dstr string
			if card.Due != "" {
				deadLine, err := time.Parse(trelloLayout, card.Due)
				if err != nil {
					return errors.Wrap(err, "faild due parse")
				}
				dstr = deadLine.Format(outputLayout)
			}

			todoList[i].Cards = append(todoList[i].Cards, TodoCard{
				Name:       card.Name,
				DeadLine:   dstr,
				Checklists: todoCheckLists,
			})
			continue
		}
	}
	return nil
}

func convertCheckLists(checklists []trello.Checklist) []TodoCheckList {
	todoCheckLists := make([]TodoCheckList, len(checklists), len(checklists))
	for i, checklist := range checklists {
		todoItems := convertTodoItems(checklist.CheckItems)
		todoCheckLists[i] = TodoCheckList{
			Name:       checklist.Name,
			CheckItems: todoItems,
		}
	}
	return todoCheckLists
}

func convertTodoItems(checkItems []trello.ChecklistItem) []TodoCheckItem {
	todoItems := make([]TodoCheckItem, len(checkItems), len(checkItems))
	for i, item := range checkItems {
		todoItems[i] = TodoCheckItem{
			Name:  item.Name,
			State: item.State,
		}
	}
	return todoItems
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
