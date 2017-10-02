package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

const format = "2006-01-02"

func outputFile(file string, output Output, outputPath string) error {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return err
	}
	now := time.Now()
	foutput, err := os.Create(fmt.Sprintf("%s-%s", now.Format(format), outputPath))
	if err != nil {
		return err
	}
	defer foutput.Close()

	return tmp.Execute(foutput, output)
}
