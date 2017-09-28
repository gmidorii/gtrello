package main

import (
	"html/template"
	"os"
)

func outputFile(file string, output Output) error {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return err
	}
	foutput, err := os.Create("./output.md")
	if err != nil {
		return err
	}
	defer foutput.Close()

	return tmp.Execute(foutput, output)
}
