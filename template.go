package main

import (
	"html/template"
	"os"
)

func outputFile(file string, output Output, outputPath string) error {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return err
	}
	foutput, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer foutput.Close()

	return tmp.Execute(foutput, output)
}
