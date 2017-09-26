package main

import (
	"flag"
)

func main() {
	config := flag.String("c", "./config.toml", "config file path")
	flag.Parse()
}
