package cmd

import (
	"flag"
	"log"
)

func Run() {
	name := flag.String("name", "", "project name")
	flag.Parse()
	log.Println(*name)
}
