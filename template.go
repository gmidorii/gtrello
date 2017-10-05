package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

const format = "2006-01-02"

func outputFile(file string, output Output, outputPath string) (string, error) {
	tmp, err := template.New("template.md").ParseFiles(file)
	if err != nil {
		return "", err
	}
	now := time.Now()
	name := fmt.Sprintf("%s/%s-%s", outputPath, now.Format(format), gtrello)
	foutput, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer foutput.Close()

	if err = tmp.Execute(foutput, output); err != nil {
		return "", err
	}
	return name, nil
}
