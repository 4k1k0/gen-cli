package main

import (
	"embed"
	"flag"
	"fmt"
	"strings"

	"github.com/4k1k0/gen-cli/internal/cli"

	justgo "github.com/gophers-mx/just-go"
)

//go:embed assets/*
var assets embed.FS

type Details struct {
	Name           string
	Age            int
	Version        string
	OriginalAuthor string
	Company        string
	URL            string
}

func main() {
	n := flag.String("name", "", "project name")
	v := flag.String("version", "", "go version")
	oa := flag.String("originalAuthor", "wako", "original author")
	c := flag.String("company", "WAKOCorp", "company")
	a := flag.Int("age", 0, "age")

	nv := cleanVersion(v)

	flag.Parse()

	generathor := &justgo.Generathor{
		Assets:      assets,
		ProjectName: n,
		Generator:   cli.NewGeneratorCLI(),
		TemplateDetails: Details{
			Name:           *n,
			Age:            *a,
			Version:        nv,
			OriginalAuthor: *oa,
			Company:        *c,
			URL:            "holi",
		},
	}

	generathor.Run()
}

func cleanVersion(version *string) (res string) {
	defer func() {
		if recover() != nil {
			res = "1.18"
		}
	}()

	vn := strings.Split(*version, "go")[1]
	vs := strings.Split(vn, ".")

	res = fmt.Sprintf("%s.%s\n", vs[0], vs[1])

	return
}
