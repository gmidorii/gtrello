package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
)

func writeFile(file, outputFile string, todo Todo, filterLists []string) error {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return err
	}
	foutput, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer foutput.Close()

	lists := make([]TodoList, len(filterLists))
	i := 0
	for _, list := range todo.Lists {
		if !containsList(list.Name, filterLists) {
			continue
		}
		lists[i] = list
		i++
	}
	filtered := Todo{
		Lists: lists,
	}

	if err = tmp.Execute(foutput, filtered); err != nil {
		return err
	}
	return nil
}

func appendFile(outputFile, outputPath string, count int) error {
	fileinfos, err := ioutil.ReadDir(outputPath)
	if err != nil {
		return err
	}

	sort.Slice(fileinfos, func(i, j int) bool {
		return fileinfos[i].ModTime().Unix() > fileinfos[j].ModTime().Unix()
	})
	var filtered []os.FileInfo
	if len(fileinfos) > count {
		filtered = fileinfos[0:count]
	} else {
		filtered = fileinfos
	}

	var text string
	for _, info := range filtered {
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", outputPath, info.Name()))
		if err != nil {
			return err
		}
		text = fmt.Sprintf("%s###%s\n%s\n", text, info.Name(), string(b))
	}

	file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, text)
	return nil
}
