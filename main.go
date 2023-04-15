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
	projectVersion := flag.String("version", "", "go version")

	flag.Parse()

	generathor := &justgo.Generathor{
		Assets:      assets,
		ProjectName: projectName,
		Version:     projectVersion,
		Generator:   cli.NewGeneratorCLI(),
	}

	generathor.Run()
}
