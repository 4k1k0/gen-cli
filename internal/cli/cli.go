package cli

import (
	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg"
)

// GeneratorCLI creates a new CLI project.
type GeneratorCLI struct{}

// NewGeneratorCLI retrieves a generator for a CLI project.
func NewGeneratorCLI() *GeneratorCLI {
	return &GeneratorCLI{}
}

// Run creates a CLI project using the embedded assets.
func (lg *GeneratorCLI) Run(cfg *config.ProjectConfig) {
	pkg.GenFromTemplate("assets", "", "README.md")
	pkg.GenFromTemplate("assets", "", "go.mod")
	pkg.GenFromTemplate("assets", "", "main.go")
	pkg.CopyDirectory("assets/cmd", "cmd")
	pkg.CopyFile("assets", "", "hello.md")
}
