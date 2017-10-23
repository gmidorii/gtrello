package main

import (
	"html/template"
	"os"
)

func writeFile(file string, output Todo, outputFile string) error {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return err
	}
	foutput, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer foutput.Close()

	if err = tmp.Execute(foutput, output); err != nil {
		return err
	}
	return nil
}
