package main

import (
	"embed"
	"flag"

	"github.com/4k1k0/gen-cli/internal/cli"

	justgo "github.com/gophers-mx/just-go"
)

//go:embed assets/*
var assets embed.FS

func main() {
	projectName := flag.String("name", "", "project name")
	projectType := flag.String("type", "", "project type (cli, etc)")

	flag.Parse()

	generator := cli.NewGeneratorCLI()

	justgo.Run(assets, projectName, projectType, generator)
}
